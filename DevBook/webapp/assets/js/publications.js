$('#new-post').on('submit', createPublication);

$(document).on('click', '.like-publication', likePublication);
$(document).on('click', '.unlike-publication', unlikePublication);

$('#update-publication').on('click', updatePublication);
$('.delete-publication').on('click', deletePublication);

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

        elementClick.addClass('unlike-publication');
        elementClick.addClass('text-danger');
        elementClick.removeClass('like-publication');

    }).fail(function () {
        alert("Error liked!");
    }).always(function() {
        elementClick.prop('disabled', false);
    });
}

function unlikePublication(evento) {
    evento.preventDefault();

    const elementClick = $(evento.target);
    const publicationID = elementClick.closest('div').data('publication-id');

    elementClick.prop('disabled', true);
    $.ajax({
        url: `/publications/${publicationID}/unlike`,
        method: "POST"
    }).done(function () {
        const counterLikes = elementClick.next('span');
        const quantityLikes = parseInt(counterLikes.text());

        counterLikes.text(quantityLikes - 1);

        elementClick.removeClass('unlike-publication');
        elementClick.removeClass('text-danger');
        elementClick.addClass('like-publication');

    }).fail(function () {
        alert("Error liked!");
    }).always(function() {
        elementClick.prop('disabled', false);
    });
}

function updatePublication(evento) {
    $(this).prop('disabled', true);

    const publicationID = $(this).data('publication-id');

    $.ajax({
        url: `/publications/${publicationID}`,
        method: "PUT",
        data: {
            title: $('#title').val(),
            content: $('#content').val()
        }
    }).done(function() {
        alert("Publication update sucessfully!");
    }).fail(function() {
        alert("Publication update error!");
    }).always(function() {
        $('#update-publication').prop('disabled', false);
    });
}

function deletePublication(evento) {
    evento.preventDefault();

    const elementClick = $(evento.target);
    const publication = elementClick.closest('div')
    const publicationID = publication.data('publication-id');

    elementClick.prop('disabled', true);

    $.ajax({
        url: `/publications/${publicationID}`,
        method: "DELETE",
    }).done(function(){
        publication.fadeOut("slow", function() {
            $(this).remove();
        });
    }).fail(function() {
        alert("Error deleting publication!")
    });
}