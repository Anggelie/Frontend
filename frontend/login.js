document.getElementById("loginForm").addEventListener("submit", async function (e) {
    e.preventDefault();

    const username = document.getElementById("username").value;
    const password = document.getElementById("password").value;

    const response = await fetch("http://localhost:3000/login", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ username, password }),
    });

    const result = await response.json();

    const msg = document.getElementById("loginMessage");
    if (result.success) {
        msg.style.color = "green";
        msg.textContent = "Inicio de sesión exitoso. Redirigiendo...";
        setTimeout(() => {
            window.location.href = `profile.html?id=${result.data.userId}`;
        }, 1500);
    } else {
        msg.style.color = "red";
        msg.textContent = ` ${result.error?.message || "Error al iniciar sesión"}`;
    }
});
