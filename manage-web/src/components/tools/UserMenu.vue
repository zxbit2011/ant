<template>
  <div class="user-wrapper">
    <div class="content-box">
      <a @click="info">
        <span class="action">
          <a-icon type="question-circle-o"></a-icon>
        </span>
      </a>
      <!--<notice-icon class="action"/>-->
      <a-dropdown>
        <span class="action ant-dropdown-link user-dropdown-menu">
          <a-avatar class="avatar" size="small" :src="avatar()"/>
          <span>{{ nickname() }}</span>
        </span>
        <a-menu slot="overlay" class="user-dropdown-menu-wrapper">
          <a-menu-item key="0">
            <router-link :to="{ name: 'PersonalBase' }">
              <a-icon type="user"/>
              <span>个人中心</span>
            </router-link>
          </a-menu-item>
          <!--<a-menu-item key="1">
            <router-link :to="{ name: 'PersonalPwd' }">
              <a-icon type="lock" />
              <span>修改密码</span>
            </router-link>
          </a-menu-item>-->
          <a-menu-item key="2">
            <router-link :to="{ name: 'PersonalCustom' }">
              <a-icon type="highlight" />
              <span>个性化</span>
            </router-link>
          </a-menu-item>
          <a-menu-divider/>
          <a-menu-item key="3">
            <a href="javascript:;" @click="handleLogout">
              <a-icon type="logout"style="color:red"/>
              <span style="color:red">退出登录</span>
            </a>
          </a-menu-item>
        </a-menu>
      </a-dropdown>
    </div>
  </div>
</template>

<script>
// import NoticeIcon from '@/components/NoticeIcon'
import { mapActions, mapGetters } from 'vuex'

export default {
  name: 'UserMenu',
  components: {
    // NoticeIcon
  },
  methods: {
    ...mapActions(['Logout']),
    ...mapGetters(['nickname', 'avatar']),

    info () {
      const h = this.$createElement
      this.$info({
        title: '使用帮助',
        content: h('div', {}, [
          h('p', ' '),
          h('p', '使用帮助')
        ]),
        onOk () {}
      })
    },
    handleLogout () {
      const that = this

      this.$confirm({
        title: '提示',
        content: '真的要注销登录吗 ?',
        onOk () {
          return that.Logout({}).then(() => {
            window.location.reload()
          }).catch(err => {
            that.$message.error({
              title: '错误',
              description: err.message
            })
          })
        },
        onCancel () {
        }
      })
    }
  }
}
</script>
