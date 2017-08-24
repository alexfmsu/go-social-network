<!DOCTYPE html>
<html>
	<head>
	    <meta charset="utf-8"/>
	    
	    <title>Index</title>

	    <script type="text/javascript" src="https://ajax.googleapis.com/ajax/libs/jquery/3.2.1/jquery.min.js"></script>
	    <script>
	    	$(document).ready(function(){
	    		$(".add_friend_form>input").click(
	    			function(){
	    				var btn = $(this);

	    				$.ajax({
  							url: "/add_friend?search=abcd10",
  							data: {user_2: $(this).attr("name")} 
						}).done(function (data) {
	    					if(data == 1){
	    						btn.val("Заявка отправлена");
	    					}else{
								alert("Some Error");
	    					}
						});	
					}
	    		);

	    		$(".confirm_friend_form>input").click(
	    			function(){
	    				$.ajax({
	  							url: "/confirm_friend",
	  							data: {user_1: $(this).attr("name")} 
							}).done(function (data) {
		    					if(data == 1){
		    						btn.val("Заявка отправлена");
		    					}else{
									alert("Some Error");
		    					}
							}
						);
					}
	    		);
	    	});
	    </script>

	    <style>
	    	.friends{
	    		width:250px;
	    		/*display:inline-block;*/
	    		float:left;
	    	}
	    	.invite{
	    		/*display:inline-block;*/
	    		width:250px;
	    		float:left;
	    	}

	    	.found_users{
	    		width:250px;
	    		float:left;
	    		border:1px solid black;

	    	}

	    	.found_users div{
	    		/*clear:both;*/
	    	}
	    </style>
	</head>

	<body>
		{{template "nav" .}}
		<div class="content">
		<p>
		
		<b>Friends</b>

		<p>
		
		<form method="GET">
			<input type="text" name="search"/>
			<input type="submit" value="Найти"/>
		</form>
		
		<p>
		{{if .friends}}
			<div style="border:1px solid black" class="friends">
			<b>Friends</b>
			{{range $k, $v := .friends}}
				<table>
					<tr>
						<td>
							login:
						</td>

						<td>
							{{$k}}
						</td>
					</tr>

					<tr>
						<td>
							email: 
						</td>
						
						<td>
						{{$v.email}}
						</td>
					</tr>

					<tr>
						<td>
							firstname: 
						</td>
						
						<td>
						{{$v.firstname}}
						</td>
					</tr>

					<tr>
						<td>
							lastname: 
						</td>
						
						<td>
						{{$v.lastname}}
						</td>
					</tr>

					<tr>
						<td>
							gender: 
						</td>
						
						<td>
						{{$v.gender}}
						</td>
					</tr>

					<tr>
						<td>
						</td>
						
						<td>
							<form class="remove_friend_form">
								<input type="button" name={{$k}} value="Remove friend"/>
							</form>
						</td>
					</tr>
				</table>
			{{end}}
			</div>
		{{end}}

		{{if .invite}}
			<div style="border:1px solid black" class="invite">
			<b>Invite</b>
			<p>

			<table>
				{{range $k, $v := .invite}}
				<tr>
						<td>
							login: 
						</td>
						
						<td>
							{{$k}}
						</td>
					</tr>

					<tr>
						<td>
							email: 
						</td>
						
						<td>
						{{$v.email}}
						</td>
					</tr>

					<tr>
						<td>
							firstname: 
						</td>
						
						<td>
						{{$v.firstname}}
						</td>
					</tr>

					<tr>
						<td>
							lastname: 
						</td>
						
						<td>
						{{$v.lastname}}
						</td>
					</tr>

					<tr>
						<td>
							gender: 
						</td>
						
						<td>
						{{$v.gender}}
						</td>
					</tr>

					<tr>
						<td>
						</td>
						
						<td>
							<form class="confirm_friend_form">
								<input type="button" name={{$k}} value="Confirm friend"/>
							</form>
						</td>
					</tr>
				{{end}}
			</table>
			</div>
		{{end}}

		{{if not .found_users}}
			Not found
		{{else}}
			<div class="found_users">
			{{range $k, $v := .found_users}}
				<div>
				
				<b>Found User</b>
				
				<p>
				
				<table>
					<tr>
						<td>
								login: 
							</td>
							
							<td>
								{{$k}}
							</td>
						</tr>

						<tr>
							<td>
								email: 
							</td>
							
							<td>
							{{$v.email}}
							</td>
						</tr>

						<tr>
							<td>
								firstname: 
							</td>
							
							<td>
							{{$v.firstname}}
							</td>
						</tr>

						<tr>
							<td>
								lastname: 
							</td>
							
							<td>
							{{$v.lastname}}
							</td>
						</tr>

						<tr>
							<td>
								gender: 
							</td>
							
							<td>
							{{$v.gender}}
							</td>
						</tr>

						<tr>
							<td>
							</td>
							
							<td>
								<form class="add_friend_form">
									<input type="button" name={{$k}} value="Add to friend"/>
								</form>
							</td>
						</tr>
					</table>
				</div>
			</div>

			<p>
			{{end}}
		{{end}}
		
		<p>
		</div>
</body>

{{template "footer" .}}

</html>