$('#form-registration').on('submit', createUser)

function createUser(evento) {
    evento.preventDefault();
    console.log("Dentro da função usuário!")

    if ($('#password').val() != $('#confirm-password').val()) {
        alert("Passwords are not equals");
        return;
    }

    $.ajax({
        url: "/users",
        method: "POST",
        data: {
           name: $('#name').val(), 
           email: $('#email').val(), 
           nick: $('#nick').val(), 
           password: $('#password').val()
        }
    }).done(function() {
        alert("User succesfully registered!");
    }).fail(function(erro) {
        console.log("Entrou no FAIL");
        console.log(erro);
        console.log("Status:", erro.status);
        console.log("Resposta:", erro.responseJSON);
        alert("Error registering user!");
    });
}