// get all the necessary DOM elements
const todoList = document.querySelector('.todo-list');
const doneList = document.querySelector('.done-list');
const addItemInput = document.querySelector('.add-item-input');
const addItemButton = document.querySelector('.add-item-button');
const deleteAllItemsButton = document.querySelector('.delete-all-items-button');

addItemButton.addEventListener('click', addItem);
addItemInput.addEventListener('keyup', (event) => {
  if (event.key === 'Enter') {
    addItem();
  }
});
deleteAllItemsButton.addEventListener('click', clearAllItems);

// initialize the lists from local storage
initializeLists();

function initializeLists() {
  const todoItems = JSON.parse(localStorage.getItem('todoItems')) || [];
  const doneItems = JSON.parse(localStorage.getItem('doneItems')) || [];

  todoItems.forEach((text) => {
    addToDoItem(text);
  });

  doneItems.forEach((text) => {
    addDoneItem(text);
  });
}

function addItem() {
  const itemText = addItemInput.value.trim();
  if (itemText !== '') {
    addToDoItem(itemText);
    addItemInput.value = '';
  }
}

function addToDoItem(text) {
  const item = document.createElement('div');
  item.classList.add('item');

  const itemText = document.createElement('span');
  itemText.classList.add('item-text');
  itemText.textContent = text;

  const itemCheckbox = document.createElement('input');
  itemCheckbox.setAttribute('type', 'checkbox');
  itemCheckbox.classList.add('item-checkbox');
  itemCheckbox.addEventListener('change', () => {
    moveItemToDoneList(item);
  });

  item.appendChild(itemText);
  item.appendChild(itemCheckbox);

  todoList.querySelector('.items').appendChild(item);
  saveListsToLocalStorage();
}

function addDoneItem(text) {
  const item = document.createElement('div');
  item.classList.add('item');

  const doneItemText = document.createElement('span');
  doneItemText.classList.add('done-item-text');
  doneItemText.textContent = text;

  const deleteButton = document.createElement('button');
  deleteButton.classList.add('delete-item-button');
  deleteButton.textContent = 'X';
  deleteButton.addEventListener('click', () => {
    deleteItem(item);
  });

  item.appendChild(doneItemText);
  item.appendChild(deleteButton);

  doneList.querySelector('.items').appendChild(item);
  saveListsToLocalStorage();
}

function moveItemToDoneList(item) {
  const itemText = item.querySelector('.item-text').textContent;
  deleteItem(item);
  addDoneItem(itemText);
}

function deleteItem(item) {
  item.remove();
  saveListsToLocalStorage();
}

function clearAllItems() {
  doneList.querySelector('.items').innerHTML = '';
  saveListsToLocalStorage();
}

function saveListsToLocalStorage() {
  const todoItems = Array.from(todoList.querySelectorAll('.item-text')).map((itemText) => itemText.textContent);
  const doneItems = Array.from(doneList.querySelectorAll('.done-item-text')).map((itemText) => itemText.textContent);

  localStorage.setItem('todoItems', JSON.stringify(todoItems));
  localStorage.setItem('doneItems', JSON.stringify(doneItems));
}