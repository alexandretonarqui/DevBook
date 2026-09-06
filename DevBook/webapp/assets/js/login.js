$('#login').on('submit', login);

function login(evento) {
    evento.preventDefault();

    $.ajax({
        url: "/login",
        method: "POST",
        data: {
            email: $('#email').val(),
            password: $('#password').val(),
        }
    }).done(function() {
        window.location = "/home";
    }).fail(function(erro) {
        Swal.fire("Ops...", "User or Password invalid!", "error");
    });
}