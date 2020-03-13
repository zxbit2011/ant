<template>
  <div class="main">
    <a-form
      v-if="debug"
      id="formLogin"
      class="user-layout-login"
      ref="formLogin"
      :form="form"
      @submit="handleSubmit"
    >
      <a-tabs
        :activeKey="customActiveKey"
        :tabBarStyle="{ textAlign: 'center', borderBottom: 'unset' }"
        @change="handleTabClick"
      >
        <a-tab-pane key="tab1" tab="账号密码登录">
          <a-form-item>
            <a-input
              size="large"
              type="text"
              placeholder="请输入登录账户"
              v-decorator="[
                'username',
                {rules: [{ required: true, message: '请输入帐户名或邮箱地址' }, { validator: handleUsernameOrEmail }], validateTrigger: 'change'}
              ]"
            >
              <a-icon slot="prefix" type="user" :style="{ color: 'rgba(0,0,0,.25)' }"/>
            </a-input>
          </a-form-item>

          <a-form-item>
            <a-input
              size="large"
              type="password"
              autocomplete="false"
              placeholder="请输入登录密码"
              v-decorator="[
                'password',
                {rules: [{ required: true, message: '请输入密码' }], validateTrigger: 'blur'}
              ]"
            >
              <a-icon slot="prefix" type="lock" :style="{ color: 'rgba(0,0,0,.25)' }"/>
            </a-input>
          </a-form-item>
          <a-form-item>
            <a-input size="large" type="text" placeholder="验证码" v-decorator="['imgCode', {rules: [{ required: true, message: '请输入图形验证码' }], validateTrigger: 'blur'}]">
              <a-icon slot="prefix" type="appstore" :style="{ color: 'rgba(0,0,0,.25)' }"/>
              <img slot="suffix" :src="imgCodeUrl" @click="updateImageCode" title="点击更换验证码" style="height: 38px;">
            </a-input>
          </a-form-item>
        </a-tab-pane>
      </a-tabs>

      <a-form-item>
        <a-checkbox v-model="rememberMe" v-decorator="['rememberMe']">记住账号</a-checkbox>
        <a
          @click="info"
          class="forge-password"
          style="float: right;"
        >忘记密码</a>
      </a-form-item>

      <a-form-item style="margin-top:24px">
        <a-button
          size="large"
          type="primary"
          htmlType="submit"
          class="login-button"
          :loading="state.loginBtn"
          :disabled="state.loginBtn"
        >确定</a-button>
      </a-form-item>

      <div class="user-login-other" v-if="false">
        <span>其他登录方式</span>
        <a>
          <a-icon class="item-icon" type="alipay-circle"></a-icon>
        </a>
        <a>
          <a-icon class="item-icon" type="taobao-circle"></a-icon>
        </a>
        <a>
          <a-icon class="item-icon" type="weibo-circle"></a-icon>
        </a>
        <router-link class="register" :to="{ name: 'register' }">注册账户</router-link>
      </div>
    </a-form>
  </div>
</template>

<script>
import pick from 'lodash.pick'
import md5 from 'md5'
import { mapActions } from 'vuex'
import { timeFix } from '@/utils/util'
import { getSmsCaptcha } from '@/api/login'
import Vue from 'vue'

export default {
  name: 'Login',
  data () {
    return {
      debug: false,
      customActiveKey: 'tab1',
      loginBtn: false,
      // login type: 0 email, 1 username, 2 telephone
      loginType: 0,
      imgCodeUrl: '/api/img/code',
      form: this.$form.createForm(this),
      state: {
        time: 60,
        loginBtn: false,
        // login type: 0 email, 1 username, 2 telephone
        loginType: 0,
        smsSendBtn: false
      },
      rememberMe: false,
      loginAccountName: '',
      username: ''
    }
  },
  created () {
    this.debug = this.$route.query.debug === 'true' || true
    if (!this.debug) {
      location.href = '/manage/sso.html'
    }
    this.username = Vue.ls.get('LoginAccountName')
    if (this.username !== '') {
      this.rememberMe = true
    }
    const { form: { setFieldsValue } } = this
    this.$nextTick(() => {
      setFieldsValue(pick(this, 'username', 'rememberMe'))
    })
  },
  methods: {
    ...mapActions(['Login', 'Logout']),
    info () {
      const h = this.$createElement
      this.$info({
        title: '使用帮助',
        content: h('div', {}, [
          h('p', ' '),
          h('p', '联系人：许老师'),
          h('p', '联系电话：18108392108')
        ]),
        onOk () {}
      })
    },
    updateImageCode () {
      this.imgCodeUrl = '/api/img/code?t=' + Date.parse(new Date())
    },
    // handler
    handleUsernameOrEmail (rule, value, callback) {
      const { state } = this
      const regex = /^([a-zA-Z0-9_-])+@([a-zA-Z0-9_-])+((\.[a-zA-Z0-9_-]{2,3}){1,2})$/
      if (regex.test(value)) {
        state.loginType = 0
      } else {
        state.loginType = 1
      }
      callback()
    },
    handleTabClick (key) {
      this.customActiveKey = key
      // this.form.resetFields()
    },
    handleSubmit (e) {
      e.preventDefault()
      const { form: { validateFields }, state, customActiveKey, Login } = this

      state.loginBtn = true

      const validateFieldsKey = customActiveKey === 'tab1' ? ['username', 'password', 'imgCode', 'rememberMe'] : ['mobile', 'captcha']
      const that = this
      validateFields(validateFieldsKey, { force: true }, (err, values) => {
        if (!err) {
          const loginParams = { ...values }
          that.loginAccountName = loginParams.username
          delete loginParams.username
          loginParams[!state.loginType ? 'email' : 'username'] = values.username
          loginParams.password = md5(values.password)
          that.rememberMe = loginParams.rememberMe || false
          Login(loginParams)
            .then(res => this.loginSuccess(res))
            .catch(err => this.requestFailed(err))
            .finally(() => {
              state.loginBtn = false
            })
        } else {
          setTimeout(() => {
            state.loginBtn = false
          }, 600)
        }
      })
    },
    getCaptcha (e) {
      e.preventDefault()
      const { form: { validateFields }, state } = this

      validateFields(['mobile'], { force: true }, (err, values) => {
        if (!err) {
          state.smsSendBtn = true

          const interval = window.setInterval(() => {
            if (state.time-- <= 0) {
              state.time = 60
              state.smsSendBtn = false
              window.clearInterval(interval)
            }
          }, 1000)

          const hide = this.$message.fileListLoading('验证码发送中..', 0)
          getSmsCaptcha({ mobile: values.mobile })
            .then(res => {
              setTimeout(hide, 2500)
              this.$notification['success']({
                message: '提示',
                description: '验证码获取成功，您的验证码为：' + res.result.captcha,
                duration: 8
              })
            })
            .catch(err => {
              setTimeout(hide, 1)
              clearInterval(interval)
              state.time = 60
              state.smsSendBtn = false
              this.requestFailed(err)
            })
        }
      })
    },
    stepCaptchaSuccess () {
      this.loginSuccess()
    },
    stepCaptchaCancel () {
      this.Logout().then(() => {
        this.loginBtn = false
        this.stepCaptchaVisible = false
      })
    },
    loginSuccess (res) {
      if (res.ret === 200) {
        if (this.rememberMe) {
          Vue.ls.set('LoginAccountName', this.loginAccountName, 1000 * 1000)
        } else {
          Vue.ls.set('LoginAccountName', '', 0)
        }
        this.$router.push({ name: 'home' })
        this.$notification.success({
          message: '欢迎',
          description: `${timeFix()}，欢迎回来`
        })
      }
    },
    requestFailed (err) {
      console.log('err', err)
      this.$notification['error']({
        message: '错误',
        description: err.msg || '请求出现错误，请稍后再试',
        duration: 4
      })
    }
  }
}
</script>

<style lang="less" scoped>
.user-layout-login {
  label {
    font-size: 14px;
  }

  .getCaptcha {
    display: block;
    width: 100%;
    height: 40px;
  }

  .forge-password {
    font-size: 14px;
  }

  button.login-button {
    padding: 0 15px;
    font-size: 16px;
    height: 40px;
    width: 100%;
  }

  .user-login-other {
    text-align: left;
    margin-top: 24px;
    line-height: 22px;

    .item-icon {
      font-size: 24px;
      color: rgba(0, 0, 0, 0.2);
      margin-left: 16px;
      vertical-align: middle;
      cursor: pointer;
      transition: color 0.3s;

      &:hover {
        color: #1890ff;
      }
    }

    .register {
      float: right;
    }
  }
}

</style>

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
      padding: 0 10px !important;
      margin-bottom: 20px !important;
    }
  }
</style>
