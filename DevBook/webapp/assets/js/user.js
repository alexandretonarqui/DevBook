$('#unfollow').on('click', unfollow);
$('#follow').on('click', follow);

function unfollow() {
    const userID = $(this).data('user-id');
    $(this).prop('disable', true);

    $.ajax({
        url: `/users/${userID}/unfollow`,
        method: "POST"
    }).done(function() {
        window.location = `/users/${userID}`;
    }).fail(function() {
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
    }).done(function() {
        window.location = `/users/${userID}`;
    }).fail(function() {
        Swal.fire("Ops...", "Error following..", "error");
        $('follow').prop('disable', false);
    });
}