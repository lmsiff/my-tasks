let notes = [];
let editingNote = null;

const API_URL = "http://localhost:8080/notes";

function loadNotes() {
  fetch(API_URL)
    .then(res => res.json())
    .then(data => {
      notes = data || [];
      renderNotes();
    })
    .catch(err => console.error("Error on loading notes:", err));
}

function renderNotes() {
  const notesList = document.getElementById('notesList');
  notesList.innerHTML = '';

  notes.forEach((note) => {
    const li = document.createElement('li');

    const noteText = document.createElement('span');
    noteText.textContent = note.content;
    noteText.className = 'note-content';

    const editBtn = document.createElement('button');
    editBtn.textContent = 'Editar';
    editBtn.className = 'edit-btn';
    editBtn.addEventListener('click', () => startEditNote(note));

    const delBtn = document.createElement('button');
    delBtn.textContent = 'Excluir';
    delBtn.className = 'delete-btn';
    delBtn.addEventListener('click', () => deleteNote(note.id));

    li.appendChild(noteText);
    li.appendChild(editBtn);
    li.appendChild(delBtn);
    notesList.appendChild(li);
  });

  resetForm();
}

function addOrUpdateNote(e) {
  const noteText = document.getElementById('noteText').value.trim();
  if (noteText === '') return;

  if (editingNote) {
    fetch(`${API_URL}/${editingNote.id}`, {
      method: "PUT",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ content: noteText })
    })
      .then(() => {
        editingNote = null;
        loadNotes();
      });
  } else {
    fetch(API_URL, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ content: noteText })
    })
      .then(res => res.json())
      .then(() => loadNotes());
  }

  document.getElementById('noteText').value = '';
}

function deleteNote(id) {
  fetch(`${API_URL}/${id}`, {
    method: "DELETE"
  }).then(() => loadNotes());
}

function startEditNote(note) {
  document.getElementById('noteText').value = note.content;
  document.getElementById('noteText').classList.add('editing');
  document.getElementById('addNoteBtn').textContent = 'Save';
  editingNote = note;
}

function resetForm() {
  document.getElementById('noteText').value = '';
  document.getElementById('noteText').classList.remove('editing');
  document.getElementById('addNoteBtn').textContent = 'Add +';
  editingNote = null;
}

window.addEventListener('DOMContentLoaded', () => {
  document.getElementById('addNoteBtn').addEventListener('click', addOrUpdateNote);
  loadNotes();
});
