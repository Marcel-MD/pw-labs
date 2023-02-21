package main

import (
	"crypto/tls"
	"fmt"
	"net"
	"net/url"
	"strings"

	"golang.org/x/net/html"
)

func get(path string) string {
	// Parse the URL
	u, err := url.Parse(path)
	if err != nil {
		fmt.Println("Error parsing URL:", err)
		return ""
	}

	// Follow redirects
	for {
		// Connect to the server
		conn, err := net.Dial("tcp", u.Host+":443")
		if err != nil {
			fmt.Println("Error connecting to server:", err)
			return ""
		}

		// Configure TLS
		tlsConn := tls.Client(conn, &tls.Config{
			ServerName: u.Hostname(),
		})

		defer tlsConn.Close()

		// Handshake with server
		err = tlsConn.Handshake()
		if err != nil {
			fmt.Println("Error during TLS handshake:", err)
			return ""
		}

		// Send the HTTP request
		fmt.Fprintf(tlsConn, "GET %s HTTP/1.1\r\nHost: %s\r\n\r\n", u.Path, u.Host)

		// Read the response
		buf := make([]byte, 1024)
		var response string
		for {
			n, err := tlsConn.Read(buf)
			if err != nil {
				break
			}
			response += string(buf[:n])
		}

		if strings.HasPrefix(response, "HTTP/1.") && (strings.Contains(response, " 301 ") || strings.Contains(response, " 302 ")) {
			// Extract the new location from the response
			location := ""
			for _, line := range strings.Split(response, "\r\n") {
				if strings.HasPrefix(line, "Location: ") {
					location = strings.TrimPrefix(line, "Location: ")
					break
				}
			}
			if location == "" {
				fmt.Println("Error following redirect: no Location header")
				return ""
			}

			// Parse the new URL
			newURL, err := url.Parse(location)
			if err != nil {
				fmt.Println("Error parsing redirect URL:", err)
				return ""
			}

			// If the new URL is relative, resolve it against the old URL
			if !newURL.IsAbs() {
				newURL = u.ResolveReference(newURL)
			}

			// Update the URL and loop again
			u = newURL
			continue
		}

		// Find the beginning of the body
		bodyIndex := strings.Index(response, "\r\n\r\n")
		if bodyIndex == -1 {
			return response
		}

		// Extract the body from the response
		body := response[bodyIndex+4:]

		head := response[:bodyIndex]

		// If the body is HTML, extract the text
		if strings.Contains(head, "text/html") {
			doc, err := html.Parse(strings.NewReader(body))
			if err != nil {
				fmt.Println("Error parsing HTML:", err)
				return ""
			}

			// Extract the text from the HTML document
			text := extractText(doc)
			// Remove extra spaces
			text = strings.Join(strings.Fields(text), " ")
			return text
		} else if strings.Contains(head, "text/plain") || strings.Contains(head, "application/json") {
			// If the body is not HTML, print it as is
			return body
		}

		return head
	}
}

// Helper function to extract text nodes from an HTML document
func extractText(n *html.Node) string {
	var text string

	switch n.Type {
	case html.ElementNode, html.DocumentNode:
		switch n.Data {
		case "script", "style", "head", "iframe", "noscript", "svg", "title":
			// Ignore these tags and their content
			return ""
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			text += extractText(c)
		}
	case html.TextNode:
		text = n.Data
	case html.CommentNode:
		return ""
	}

	return text
}
