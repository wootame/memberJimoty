<template>
  <div id="profile" class="main-stage">
    <h2>プロフィール</h2>
    <p><span>名前: </span>{{ user.name }}</p>
    <p><span>LoginID: </span>{{ user.login_id }}</p>
    <p><span>Email: </span>{{ user.email }}</p>
    <div class="content">
      <span class="title">自己紹介</span><br>
      {{ user.profile }}
    </div>
    <router-link to="/user/edit/1">
      <button>編集</button>
    </router-link>
  </div>
</template>

<script>
import EventListComponents from './EventList.vue';

export default {
  components: {
    EventListComponents,
  },
  data() {
    return {
      user: {
        name: 'John Doe',
        login_id: 'LoginID',
        email: 'john@example.com',
        profile: 'プロフィールです。プロフィールです。プロフィールです。',
      },
      events: [
        { 
          id: 1, 
          title: 'イベント 1' ,
          category: '飲み会',
          atmosphere: 'しっぽり',
          tags: ['飲み会', 'しっぽり', '友達づくり'],
          description: 'こんなイベントです！',
        },
        { 
          id: 2, 
          title: 'イベント 2',
          category: 'スポーツ',
          atmosphere: 'わいわい',
          tags: ['スポーツ', '健康意識', 'アクティブ'],
          description: 'こんなイベントです', 
        },
        { 
          id: 3, 
          title: 'イベント 1' ,
          category: '飲み会',
          atmosphere: 'ワイワイ',
          tags: ['飲み会', '楽しく', '大勢'],
          description: 'こんなイベントです！',
        },
      ],
    };
  },
};
</script>

<style scoped>
#profile {
  border: solid 1px #888;
  border-radius: 5px;
  padding: 0 20px 20px;
}

#profile p {
  font-size: 18px;
}

#profile > span {
  width: 80px;
  display: inline-block;
}

.content {
  margin-top: 20px;
}

#profile .title {
  font-size: 18px;
  margin-bottom: 10px;
}

button {
  width: 100px;
  text-align: center;
  margin: 20px auto;
  display: block;
}
</style>
