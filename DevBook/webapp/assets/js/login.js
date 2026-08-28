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
        console.log("Entrou no FAIL");
        console.log(erro);
        console.log("Status:", erro.status);
        console.log("Resposta:", erro.responseJSON);
        alert("User or Password invalid!");
    });
}