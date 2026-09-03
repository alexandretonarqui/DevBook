$('#new-post').on('submit', createPublication)
$('.like-publication').on('click', likePublication);

function createPublication(evento) {
    evento.preventDefault();

    $.ajax({
        url: "/publications",
        method: "POST",
        data: {
            title: $('#title').val(),
            content: $('#content').val(),
        }
    }).done(function () {
        window.location = "/home";
    }).fail(function () {
        alert("Error creating publication!");
    })
}

function likePublication(evento) {
    evento.preventDefault();

    const elementClick = $(evento.target);
    const publicationID = elementClick.closest('div').data('publication-id');

    elementClick.prop('disabled', true);
    $.ajax({
        url: `/publications/${publicationID}/like`,
        method: "POST"
    }).done(function () {
        const counterLikes = elementClick.next('span');
        const quantityLikes = parseInt(counterLikes.text());

        counterLikes.text(quantityLikes + 1);
    }).fail(function () {
        alert("Error liked!");
    }).always(function() {
        elementClick.prop('disabled', false);
    });
}