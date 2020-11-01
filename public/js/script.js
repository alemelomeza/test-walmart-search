document.addEventListener("DOMContentLoaded", function() {
    document.querySelectorAll(".card-currency").forEach(function(c) {
        c.innerHTML = new Intl.NumberFormat().format(parseInt(c.innerHTML));
    });
});