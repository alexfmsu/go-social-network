<!DOCTYPE html>
<html>
	{{template "header" .}}

	<body>
		<form method="POST">
			<table>
				<tr>
					<td>
						Login			
					</td>

					<td>
						<input type="text" name="username"/>	
					</td>
				</tr>

				<tr>
					<td>
						Password			
					</td>

					<td>
						<input type="password" name="password"/>
					</td>
				</tr>

				<tr>
					<td>
					</td>

					<td>
        				<input type="submit" value="Login"/>
					</td>
				</tr>
			</table>
		</form>
	</body>
	
	{{template "footer"}}
</html>