$('#unfollow').on('click', unfollow);
$('#follow').on('click', follow);
$('#edit-user').on('submit', edit);

function unfollow() {
    const userID = $(this).data('user-id');
    $(this).prop('disable', true);

    $.ajax({
        url: `/users/${userID}/unfollow`,
        method: "POST"
    }).done(function () {
        window.location = `/users/${userID}`;
    }).fail(function () {
        Swal.fire("Ops...", "Error unfollowing..", "error");
        $('unfollow').prop('disable', false);
    });
}

function follow() {
    const userID = $(this).data('user-id');
    $(this).prop('disable', true);

    $.ajax({
        url: `/users/${userID}/follow`,
        method: "POST"
    }).done(function () {
        window.location = `/users/${userID}`;
    }).fail(function () {
        Swal.fire("Ops...", "Error following..", "error");
        $('follow').prop('disable', false);
    });
}

function edit(evento) {
    evento.preventDefault();

    $.ajax({
        url: "/edit-user",
        method: "PUT",
        data: {
            name: $('#name').val(),
            email: $('#email').val(),
            nick: $('#nick').val(),
        }
    }).done(function() {
        Swal.fire("Success!", "Profile Updated!", "success")
            .then(function() {
                window.location = "/profile";
            });
    }).fail(function() {
        Swal.fire("Ops..", "Error Updating Profile!", "error");
    });
}