<template>
  <div>
    <a-spin v-if="loginFlag==='-1'" tip="正在验证登录信息。。。">
      <div class="spin-content" style="height: 100px">
      </div>
    </a-spin>
    <a-card v-else-if="loginFlag==='2'" :bordered="false" style="margin: -24px -24px 0px;">
      <result type="success" :title="successTitle">
        <template slot="action">
          <a-button type="primary" @click="goHome">进入主页</a-button>
          <a-button style="margin-left: 8px" @click="$router.push({ name: 'home'})" v-if="isRole">进入管理中心</a-button>
        </template>
        <div>
          <div style="font-size: 16px; color: rgba(0, 0, 0, 0.85); font-weight: 500; margin-bottom: 16px">
          </div>
          <div style="margin-bottom: 16px">
            <a-icon type="check" style="color:#52c41a" />
            {{ msg }}
          </div>
        </div>
      </result>
    </a-card>
    <a-card v-else-if="loginFlag==='0'" :bordered="false" style="margin: -24px -24px 0px;">
      <result type="error" description="请核对以下信息后，再重新登录。" title="登录失败">
        <template slot="action">
          <a href="/manage/sso.html"> <a-button type="primary" >重新登录</a-button></a>
        </template>
        <div>
          <div style="font-size: 16px; color: rgba(0, 0, 0, 0.85); font-weight: 500; margin-bottom: 16px">
          </div>
          <div style="margin-bottom: 16px">
            <a-icon type="close-circle-o" style="color: #f5222d; margin-right: 8px"/>
            {{ msg }}
          </div>
        </div>
      </result>
    </a-card>
  </div>
</template>

<script>
import { Result } from '@/components'
import { mapActions } from 'vuex'
import store from '../../store'
import { ACCESS_TOKEN } from '@/store/mutation-types'
import Vue from 'vue'

var that
export default {
  components: {
    Result
  },
  data () {
    return {
      webUrl: '/web/index.html',
      loginFlag: '0',
      msg: '',
      successTitle: '登录成功',
      UserType: null,
      isRole: false
    }
  },
  created () {
    that = this
    that.loginFlag = that.$route.query.state || '0'
    that.msg = that.$route.query.msg || '登录失败'
    if (that.loginFlag === '2') {
      const token = that.$route.query.token || null
      const tokenExpire = that.$route.query.tokenExpire || 0
      const name = that.$route.query.name || ''
      that.successTitle = `您好！${name}，登录成功`
      Vue.ls.set(ACCESS_TOKEN, token, parseInt(tokenExpire) * 1000)
      store.commit('SET_TOKEN', token)
      that.isRole = that.$route.query.r === 'true'
    }
  },
  mounted () {
    if (this.isMobile()) {
      this.webUrl = '/web/mobile/index.html'
    }
  },
  methods: {
    ...mapActions(['LoginCheck']),
    isMobile () {
      const flag = navigator.userAgent.match(/(phone|pad|pod|iPhone|iPod|ios|iPad|Android|Mobile|BlackBerry|IEMobile|MQQBrowser|JUC|Fennec|wOSBrowser|BrowserNG|WebOS|Symbian|Windows Phone)/i)
      return flag
    },
    loginCheck () {
      const {
        LoginCheck
      } = this
      LoginCheck(this.$route.query).then(res => {
        that.msg = res.msg
        if (res.ret === 200) {
          const result = res.data
          const loginInfo = result.loginInfo
          that.successTitle = `您好！${loginInfo.Name}，登录成功`
          that.loginFlag = 2
          that.UserType = loginInfo.UserType
          if (loginInfo.IsSuperAdmin || (loginInfo.SysMenu !== null && loginInfo.SysMenuBtns !== null)) {
            that.isRole = true
          }
          that.IsSuperAdmin = loginInfo.IsSuperAdmin
          that.isMenus = loginInfo.SysMenu !== null
          that.isMenuBtns = loginInfo.SysMenuBtns !== null
        } else {
          that.loginFlag = 3
        }
      })
    },
    goHome () {
      location.href = this.webUrl
    }
  }
}
</script>
<style>

  /* 手机等小屏幕手持设备 */
  @media screen and (min-width: 320px) and (max-width: 480px) {
    #userLayout.user-layout-wrapper .container{
      padding-top: 22px !important;
    }
    #userLayout.user-layout-wrapper .container .top .header .title {
      font-size: 18px !important;
    }
    #userLayout.user-layout-wrapper .container .top .header .logo{
      height: 30px !important;
      margin-top: 8px !important;
      margin-right: 8px !important;
    }
    #userLayout.user-layout-wrapper .container .top .desc{
      font-size: 12px !important;
      padding: 0 10px;
    }
    .result .icon{
      font-size: 52px !important;
      line-height: 52px !important;
      margin-bottom: 10px !important;
    }
    .ant-card{
      margin: 0 !important;
    }
    .ant-card-body{
      padding:20px 0 !important;
    }
  }
</style>
