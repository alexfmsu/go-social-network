<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Edit profile</title>
</head>
<body>

<form method="POST" action="">
    Username: <input type="text" name="username" value="{{.Login}}" /><br />
    Email: <input type="email" name="email" value="{{.Email}}"  /><br />
    First Name: <input type="text" name="first_name" value="{{.First_name}}"  /><br />
    Last Name: <input type="text" name="last_name" value="{{.Last_name}}"  /><br />
    <br>
    <br>
    Change password (if no, set empty) <br>
    Password: <input type="password" name="password"  /><br />
    New password: <input type="password" name="new_password"  /><br />
    <p><input type="submit" value="Submit" /></p>
</form>



</br>

</body>


</body>
</html>