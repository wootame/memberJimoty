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
      user: {
        name: 'John Doe',
        email: 'john@example.com',
        // Add more user details as needed
      },
    };
  },
  methods: {
    saveProfile() {
      // Handle profile save logic
      console.log('Profile saved:', this.user);

      // リダイレクト
      const profileId = 1; // プロファイルのID
      this.$router.push(`/user/profile/${profileId}`);
    },
  },
};
</script>

<style>
  .profile-form {
    width: 80%;
    border: 1px solid #888;
    border-radius: 5px;
    padding: 0 20px 20px;
    margin: 20px auto;
  }

  .form-group {
    margin: 20px 0;
  }

  .form-group label {
    font-size: 22px;
    display: block;
  }

  .form-group input[type="text"],
  .form-group input[type="email"],
  .form-group textarea,
  select {
    width: 80%;
    font-size: 20px;
    border: solid 1px #888;
    border-radius: 5px;
    margin-left: 8%;
  }

  .event-tags {
    margin-top: 10px;
    margin-left: 8%;
  }

  .event-tags span {
    background-color: #888;
    font-size: 18px;
    color: #fff;
    border-radius: 5px;
    padding: 5px;
  }

  .event-tags span + span {
    margin-left: 10px;
  }

  .form-group input[type="text"],
  .form-group input[type="email"],
  select {
    height: 40px;
  }

  #tags {
    width: 70%;
  }

  .form-group button {
    width: 6%;
    text-align: center;
    margin-left: 2%;
  }

  .form-group textarea {
    min-height: 200px;
  }

  input[type="submit"] {
    background-color: #888;
    width: 200px;
    height: 40px;
    text-align: center;
    color: #fff;
    border-radius: 5px;
    margin: 0 auto;
    display: block;
  }
</style>
