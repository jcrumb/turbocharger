$(document).ready(function(){
	$('#orderform').on('submit', function(e) {
		if ($('#usernameInput').val() == '') {
			alert("Please enter your social club name");
		}
		else {
			var formObject = $('#orderform').serializeJSON();
			console.log(formObject)
			e.preventDefault();
			$.ajax({
				url: 'http://turbo.jaycrumb.me:8227/order',
				cache: false,
				type: 'POST',
				data: JSON.stringify(formObject),
				success: function(json) {
					$('#successmessage').addClass("in");
					$('#successmessage').show();
				},
				error: function(json) {
					$('#errmessage').addClass("in");
					$('#errmessage').show();
				}
			});
		}
	});

	$('.dismiss-button').on('click', function(e) {
		$(this).parent().hide();
	});
});
