<template>
  <div id="chatList" class="main-stage">
    <h2>メッセージ 一覧</h2>
    <div>
      <div v-for="chat in chatList" :key="chat.id" class="chat">
        <router-link :to="`/chat/${chat.id}`">
          <h3><img src="{{ chat.userImagePath }}">{{ chat.userName }}</h3>
        </router-link>
        <hr>
        <p class="last-chat">{{ chat.lastChat }}</p>
        <p class="last-chat-date">{{ chat.lastChatDate }}</p>
      </div>
    </div>
  </div>
</template>

<script>
import HeaderComponent from './template/Header.vue';
export default {
  components: {
    HeaderComponent,
  },
  data() {
    return {
      chatList: [
        { 
          userId: 1, 
          userName: 'タモリ',
          userImagePath: 'https://google.com',
          lastChat: 'イベント 1についてどう思いますか。' ,
          lastChatDate: '2023/11/01 12:01',
        },
        { 
          userId: 2, 
          userName: 'モリタ',
          userImagePath: 'https://google.com',
          lastChat: 'イベント 2についてどう思いますか。' ,
          lastChatDate: '2023/11/01 12:01',
        },
        { 
          userId: 3, 
          userName: 'タモリ',
          userImagePath: 'https://google.com',
          lastChat: 'イベント 3についてどう思いますか。' ,
          lastChatDate: '2023/11/01 12:01',
        },
      ],
    };
  },
};
</script>

<style>
    .last-chat {
        color: #888;
        font-size: 16px !important;
    }

    .last-chat-date {
        font-size: 14px !important;
    }
</style>
