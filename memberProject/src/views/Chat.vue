<template>
  <div id="profile-form" class="main-stage">
    <h2>プロフィール 編集</h2>
    <form @submit.prevent="saveProfile">
      <div class="form-group">
        <label>名前: </label>
        <input type="text" v-model="user.name" />
      </div>
      <div class="form-group">
        <label for="">LoginID: </label>
        <input type="text">
      </div>
      <div class="form-group">
        <label>Email: </label>
        <input type="email" v-model="user.email" />
      </div>
      <div class="form-group">
        <label for="">自己紹介: </label>
        <textarea name="" id="" cols="30" rows="10"></textarea>
      </div>
      <input type="submit" value="変更" @click="saveProfile">
    </form>
  </div>
</template>

<script>
export default {
  data() {
    return {
      userId: 1,
      userName: 'タモリ',
      chat: {
        chat: [
          {
            send: 1,
            content: 'content, content',
            date: "2023/11/02 10:12",
          }
        ]
      }
    }
  }
};
</script>

<style>
  .chat {
      width: 80%;
  }

  .chat-form {
    width: 80%;
    border: 1px solid #888;
    border-radius: 5px;
    padding: 0 20px 20px;
    margin: 20px auto;
  }
</style>
