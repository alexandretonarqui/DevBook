$('#unfollow').on('click', unfollow);
$('#follow').on('click', follow);
$('#edit-user').on('submit', edit);
$('#updatepassword').on('submit', updatePassword);
$('#delete-user').on('click', deleteUser);

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
    }).done(function () {
        Swal.fire("Success!", "Profile Updated!", "success")
            .then(function () {
                window.location = "/profile";
            });
    }).fail(function () {
        Swal.fire("Ops..", "Error Updating Profile!", "error");
    });
}

function updatePassword(evento) {
    evento.preventDefault();

    if ($('#new-password').val() != $('#confirm-password').val()) {
        Swal.fire("Ops...", "Passwords don't match!", "warning");
        return;
    }

    $.ajax({
        url: "/updatepassword",
        method: "POST",
        data: {
            current: $('#current-password').val(),
            new: $('#new-password').val()
        }
    }).done(function () {
        Swal.fire("Success", "Passwords Updated Successfully!", "success")
            .then(function () {
                window.location = "/profile";
            })
    }).fail(function () {
        Swal.fire("Ops...", "Error Updating Password!", "error");
    });
}

function deleteUser() {
    Swal.fire({
        title: "Atention!",
        text: "Are You Sure?",
        showCancelButton: true,
        cancelButtonText: "Cancel",
        icon: "warning"
    }).then(function(confirmation) {
        if (confirmation.value) {
            $.ajax({
                url: "/delete-user",
                method: "DELETE"
            }).done(function() {
                Swal.fire("Success!", "Your Profile Deleted Successfully!", "success")
                    .then(function() {
                        window.location = "/logout";
                    })
            }).fail(function() {
                Swal.fire("Ops..", "Error Deleting User's Profile!", "error");
            });
        }
    })
}