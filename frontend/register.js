document.getElementById("registerForm").addEventListener("submit", async function (e) {
    e.preventDefault();

    const username = document.getElementById("username").value;
    const password = document.getElementById("password").value;

    const response = await fetch("http://localhost:3000/register", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ username, password }),
    });

    const result = await response.json();

    const msg = document.getElementById("registerMessage");
    if (result.success) {
        msg.style.color = "green";
        msg.textContent = " Registro exitoso, ahora inicia sesión.";
        setTimeout(() => {
            window.location.href = "login.html";
        }, 1500);
    } else {
        msg.style.color = "red";
        msg.textContent = ` ${result.error?.message || "Error al registrar"}`;
    }
});
