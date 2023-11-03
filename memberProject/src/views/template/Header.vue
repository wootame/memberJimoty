<template>
    <header>
        <a href="/">
            <img src="/public/jimoty_logo.jpeg" alt="サンプル画像">
        </a>
        <ul>
            <li><router-link to="/">ホーム</router-link></li>
            <li><router-link to="/events/edit">イベント作成</router-link></li>
            <li><router-link to="/user/profile/1">プロフィール</router-link></li>
            <li><router-link to="/chat">チャット</router-link></li>
            <li><router-link to="/logout">ログアウト</router-link></li>
        </ul>
    </header>
</template>

<script>
</script>

<style>
    header {
        display: flex;
        padding: 20px;    
    }

    header img {
        width: 200px;
    }
</style>
