{{define "nav"}}
<style>
	nav{
		display:inline-block;
		width:100%;

		padding: 0 0 30px 0;
	}
	nav a{
		width:150px;
		float:left;
		text-decoration: none;
		color:#FFFFFF;
		background:#2B3137;
		text-align:center;
		padding:10px 0 10px 0;
		//border:1px solid black;
	}
</style>

<nav>
	<a href="/profile">
		Профиль
	</a>

	<a href="/friends">
		Друзья
	</a>

	<a href="/messages">
		Сообщения
	</a>

	<a href="/photos">
		Фотографии
	</a>

	<a href="/video">
		Видеозаписи
	</a>
	
	<a href="/audio">
		Музыка
	</a>

	<a href="/logout">
		Выйти
	</a>
</nav>
{{end}}
