<!DOCTYPE html>
<html>
	{{template "header" .}}

	<body>
		<div class="content">
		{{template "nav" .}}

		<b>Profile</b>
		
		<p>

		<table>
			<tr>
				<td>
				</td>

				<td>
					<span style="color:green">@{{.Login}}</span>	
				</td>
			</tr>

			<tr>
				<td>
					LastName
				</td>

				<td>
					{{.Last_name}} {{.First_name}}	
				</td>
			</tr>

			<tr>
				<td>
					Email
				</td>

				<td>
					{{.Email}}	
				</td>
			</tr>

			<tr>
				<td>
					Пол:
				</td>

				<td>
					{{.Gender}}	
				</td>
			</tr>
		</table>

		<p>
		
		<form enctype="multipart/form-data" method="post">
		    <input type="file" name="uploadfile" />
		    <!-- <input type="hidden" name="token" value="{{.}}"/> -->
		    <input type="submit" value="upload" />
		</form>

		NewFriends: +{{.NewFriends}}
		
		<p>
		
		NewMessages: +{{.NewMessages}}
		<p>

		<a href="/edit">Edit profile</a>

		<p>
		</div>
	</body>

	{{template "footer" .}}
</html>
