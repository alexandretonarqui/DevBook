$('#form-registration').on('submit', createUser)

function createUser(evento) {
    evento.preventDefault();
    console.log("Dentro da função usuário!")

    if ($('#password').val() != $('#confirm-password').val()) {
        Swal.fire("Ops...", "Passwords are not equals", "error");
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
        Swal.fire("Success", "User succesfully registered!", "success")
            .then(function() {
                $.ajax({
                    url: "/login",
                    method: "POST",
                    data: {
                        email: $('#email').val(),
                        password: $('#password').val()
                    }
                }).done(function() {
                    window.location = "/home";
                }).fail(function() {
                    Swal.fire("Ops...", "Error authenticate user!", "error");
                })
            })
    }).fail(function(erro) {
        console.log("Entrou no FAIL");
        console.log(erro);
        console.log("Status:", erro.status);
        console.log("Resposta:", erro.responseJSON);
        Swal.fire("Ops...", "Error registering user!", "error");
    });
}