$('#new-post').on('submit', createPublication)

function createPublication(evento) {
    evento.preventDefault();

    $.ajax({
        url: "/publications",
        method: "POST",
        data: {
            title: $('#title').val(),
            content: $('#content').val(),
        }
    }).done(function() {
        window.location = "/home";
    }).fail(function() {
        alert("Error creating publication!");
    })
}