const USER_API = "http://localhost:8081";
const PRODUCTIVITY_API = "http://localhost:8080";

function login() {
    const email = document.getElementById("email").value;
    const password = document.getElementById("password").value;
    fetch(`${USER_API}/login`, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ email, password })
    })
        .then(res => res.json())
        .then(data => {
            localStorage.setItem("token", data.token);
            window.location.href = "dashboard.html";
        })
        .catch(() => alert("Login failed"));
}

function register() {
    const email = document.getElementById("regEmail").value;
    const password = document.getElementById("regPassword").value;
    fetch(`${USER_API}/register`, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ email, password })
    })
        .then(res => res.json())
        .then(() => alert("Registered! Now login."))
        .catch(() => alert("Register failed"));
}

function addTodo() {
    const text = document.getElementById("newTodo").value;
    fetch(`${PRODUCTIVITY_API}/todos`, {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
            "Authorization": `Bearer ${localStorage.getItem("token")}`
        },
        body: JSON.stringify({ text })
    }).then(() => loadTodos());
}

function loadTodos() {
    fetch(`${PRODUCTIVITY_API}/todos`, {
        headers: {
            "Authorization": `Bearer ${localStorage.getItem("token")}`
        }
    })
        .then(res => res.json())
        .then(todos => {
            const list = document.getElementById("todoList");
            list.innerHTML = "";
            todos.forEach(todo => {
                const li = document.createElement("li");
                li.textContent = todo.text;
                const btn = document.createElement("button");
                btn.textContent = "Delete";
                btn.onclick = () => deleteTodo(todo.ID);
                li.appendChild(btn);
                list.appendChild(li);
            });
        });
}

function deleteTodo(id) {
    fetch(`${PRODUCTIVITY_API}/todos/${id}`, {
        method: "DELETE",
        headers: {
            "Authorization": `Bearer ${localStorage.getItem("token")}`
        }
    }).then(() => loadTodos());
}

function logout() {
    localStorage.removeItem("token");
    window.location.href = "index.html";
}

if (window.location.pathname.includes("dashboard.html")) {
    loadTodos();
}
