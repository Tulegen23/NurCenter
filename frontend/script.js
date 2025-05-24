const token = localStorage.getItem("token");
if (!token) {
  window.location.href = "index.html"; // перенаправление, если токена нет
}

async function fetchTodos() {
  try {
    const response = await fetch("http://localhost:8080/todos", {
      headers: {
        Authorization: `Bearer ${token}`,
      },
    });

    if (!response.ok) {
      throw new Error(`Ошибка загрузки: ${response.status}`);
    }

    const data = await response.json();

    // Если у API структура { todos: [...] }, то использовать data.todos
    const todos = data.todos || data;

    const list = document.getElementById("todoList");
    list.innerHTML = "";

    todos.forEach((todo) => {
      const li = document.createElement("li");
      li.textContent = todo.title + (todo.isDone ? " ✅" : "");
      list.appendChild(li);
    });
  } catch (err) {
    console.error(err);
    alert("Не удалось загрузить задачи. Попробуйте позже.");
  }
}

async function addTodo() {
  const titleInput = document.getElementById("newTodo");
  const title = titleInput.value.trim();

  if (!title) {
    alert("Введите название задачи");
    return;
  }

  try {
    const response = await fetch("http://localhost:8080/todos", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        Authorization: `Bearer ${token}`,
      },
      body: JSON.stringify({ title }),
    });

    if (!response.ok) {
      throw new Error(`Ошибка добавления: ${response.status}`);
    }

    titleInput.value = "";
    await fetchTodos();
  } catch (err) {
    console.error(err);
    alert("Не удалось добавить задачу. Попробуйте позже.");
  }
}

function logout() {
  localStorage.removeItem("token");
  window.location.href = "index.html";
}

document.getElementById("addBtn").addEventListener("click", addTodo);
window.addEventListener("load", fetchTodos);
