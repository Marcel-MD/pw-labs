package main

import (
	"crypto/tls"
	"fmt"
	"net"
	"net/url"
	"regexp"
	"strings"
	"time"

	"golang.org/x/net/html"
)

func get(path string) (string, string) {

	path = addHttpsPrefix(path)

	// Parse the URL
	u, err := url.Parse(path)
	if err != nil {
		fmt.Println("Error parsing URL:", err)
		return "", ""
	}

	// Follow redirects
	for {
		// Connect to the server
		conn, err := net.Dial("tcp", u.Host+":443")
		if err != nil {
			fmt.Println("Error connecting to server:", err)
			return "", ""
		}

		// Configure TLS
		tlsConn := tls.Client(conn, &tls.Config{
			ServerName: u.Hostname(),
		})
		// Read timeout
		tlsConn.SetReadDeadline(time.Now().Add(5 * time.Second))

		defer tlsConn.Close()

		// Handshake with server
		err = tlsConn.Handshake()
		if err != nil {
			fmt.Println("Error during TLS handshake:", err)
			return "", ""
		}

		// Send the HTTP request
		fmt.Fprintf(tlsConn, "GET %s HTTP/1.1\r\nHost: %s\r\nUser-Agent: %s\r\nAccept: %s\r\n\r\n", u.RequestURI(), u.Host, "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Safari/537.36", "text/html,text/plain;q=0.9,application/json;q=0.8,*/*;q=0.7")

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
				return "", ""
			}

			// Parse the new URL
			newURL, err := url.Parse(location)
			if err != nil {
				fmt.Println("Error parsing redirect URL:", err)
				return "", ""
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
			fmt.Println(response)
		}

		// Extract the body from the response
		body := response[bodyIndex+4:]

		head := response[:bodyIndex]

		// If the body is HTML, extract the text
		if strings.Contains(head, "text/html") {
			return body, "text/html"
		} else if strings.Contains(head, "text/plain") {
			return body, "text/plain"
		} else if strings.Contains(head, "application/json") {
			return body, "application/json"
		} else {
			return "", ""
		}
	}
}

func processBody(body string, bodyType string) {
	switch bodyType {
	case "text/html":
		doc, err := html.Parse(strings.NewReader(body))
		if err != nil {
			fmt.Println("Error parsing HTML:", err)
			return
		}

		// Extract the text from the HTML document
		text := extractText(doc)
		text = prettyText(text)
		fmt.Println(text)
	case "text/plain", "application/json":
		// If the body is not HTML, print it as is
		fmt.Println(body)
	}
}

func extractText(n *html.Node) string {
	var text string

	switch n.Type {
	case html.ElementNode, html.DocumentNode:
		switch n.Data {
		case "script", "style", "head", "iframe", "noscript", "svg", "title":
			// Ignore these tags and their content
			return ""

		case "br", "p", "div", "li":
			// Add a newline after these tags
			text = "\n"
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

func prettyText(text string) string {
	// Replace consecutive whitespace characters with a single space
	spaceRegex := regexp.MustCompile(`[ \t\r\f\v]+`)
	text = spaceRegex.ReplaceAllString(text, " ")

	// Replace consecutive newlines (which might have one space in between) with a single newline
	newlineRegex := regexp.MustCompile(`(\n[\p{Zs}\t]*){2,}`)
	text = newlineRegex.ReplaceAllString(text, "\n")

	// Trim leading and trailing whitespace
	text = strings.TrimSpace(text)

	return text
}

func addHttpsPrefix(url string) string {
	if !strings.HasPrefix(url, "https://") {
		return "https://" + url
	}
	return url
}
