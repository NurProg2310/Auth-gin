(() => {
  const form = document.getElementById("loginForm");
  const email = document.getElementById("email");
  const password = document.getElementById("password");

  const emailError = document.getElementById("emailError");
  const passwordError = document.getElementById("passwordError");

  const successMessage = document.getElementById("successMessage");
  const loginBtn = document.querySelector(".login-btn");
  const passwordToggle = document.getElementById("passwordToggle");

  const wrapperOf = (input) => input.closest(".input-wrapper");

  function setError(input, errorEl, message) {
    const w = wrapperOf(input);
    if (w) w.classList.add("invalid");
    errorEl.textContent = message || "";
  }

  function clearError(input, errorEl) {
    const w = wrapperOf(input);
    if (w) w.classList.remove("invalid");
    errorEl.textContent = "";
  }

  function isValidEmail(v) {
    return /^[^\s@]+@[^\s@]+\.[^\s@]{2,}$/i.test(v.trim());
  }

  function validate() {
    let ok = true;

    const ev = email.value.trim();
    const pv = password.value;

    if (!ev) { setError(email, emailError, "Email is required"); ok = false; }
    else if (!isValidEmail(ev)) { setError(email, emailError, "Enter a valid email"); ok = false; }
    else { clearError(email, emailError); }

    if (!pv) { setError(password, passwordError, "Password is required"); ok = false; }
    else { clearError(password, passwordError); }

    return ok;
  }

  if (passwordToggle) {
    passwordToggle.addEventListener("click", () => {
      password.type = (password.type === "password") ? "text" : "password";
    });
  }

  email.addEventListener("input", () => clearError(email, emailError));
  password.addEventListener("input", () => clearError(password, passwordError));

  form.addEventListener("submit", async (e) => {
    e.preventDefault();
    if (!validate()) return;

    loginBtn.classList.add("loading");
    loginBtn.disabled = true;

    try {
      const res = await fetch("/api/login", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({
          email: email.value.trim(),
          password: password.value
        })
      });

      const data = await res.json().catch(() => ({}));

      if (!res.ok || !data.ok) {
        setError(password, passwordError, data.error || "Wrong email or password");
        return;
      }

      successMessage?.classList.add("show");

      // пример: редирект на /dashboard через 1 сек (когда сделаешь страницу)
      // setTimeout(() => window.location.href = "/dashboard", 1000);

      setTimeout(() => successMessage?.classList.remove("show"), 1500);

    } catch (err) {
      setError(password, passwordError, "Server error");
    } finally {
      loginBtn.classList.remove("loading");
      loginBtn.disabled = false;
    }
  });
})();
