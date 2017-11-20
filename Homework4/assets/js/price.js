$(document).ready(function() {
    $.ajax({
        url: "/api/test"
    }).then(function(x) {
       $('.item-name').append(x.name);
       $('.item-price').append(x.price);
    });
});
