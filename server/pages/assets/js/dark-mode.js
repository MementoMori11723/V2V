function toggleDarkMode() {
  document.body.classList.toggle("dark-mode"),
    localStorage.setItem(
      "darkMode",
      document.body.classList.contains("dark-mode"),
    ),
    document
      .querySelectorAll(".toggle-switch input[type='checkbox']")
      .forEach((e) => {
        e.checked = document.body.classList.contains("dark-mode");
      });
}
function toggleSidebar() {
  let e = document.querySelector(".sidebar-overlay"),
    o = document.querySelector(".sidebar"),
    t = document.querySelector(".hamburger");
  e.classList.toggle("show"),
    o.classList.toggle("show"),
    t.classList.toggle("open");
}
function closeSidebar() {
  let e = document.querySelector(".sidebar-overlay"),
    o = document.querySelector(".sidebar"),
    t = document.querySelector(".hamburger");
  e.classList.remove("show"),
    o.classList.remove("show"),
    t.classList.remove("open");
}
document.addEventListener("DOMContentLoaded", () => {
  let e = "true" === localStorage.getItem("darkMode");
  e && document.body.classList.add("dark-mode"),
    document
      .querySelectorAll(".toggle-switch input[type='checkbox']")
      .forEach((o) => {
        o.checked = e;
      });
});
