<template>
  <a-modal
    title="操作权限编辑"
    :width="1000"
    :visible="visible"
    :keyboard="false"
    :maskClosable="false"
    :destroyOnClose="true"
    :confirmLoading="confirmLoading"
    @cancel="handleCancel"
  >
    <a-spin :spinning="confirmLoading">
      <a-form :form="form">
        <!-- step1 -->
        <div>
          <a-form-item
            label="按钮名称"
            :labelCol="labelCol"
            :wrapperCol="wrapperCol">
            <a-input v-decorator="['name', {rules: [{required: true}]}]" placeholder="请输入按钮名称" />
          </a-form-item>
          <a-form-item
            label="权限标识"
            :labelCol="labelCol"
            :wrapperCol="wrapperCol">
            <a-input v-decorator="['permission', {rules: [{required: true}]}]" placeholder="请输入权限标识" />
          </a-form-item>
          <a-form-item
            label="请求方式"
            :labelCol="labelCol"
            :wrapperCol="wrapperCol" >
            <a-radio-group
              v-decorator="['method', {rules: [{required: true}]}]">
              <a-radio value="POST">POST</a-radio>
              <a-radio value="GET">GET</a-radio>
            </a-radio-group>
          </a-form-item>
          <a-form-item
            label="请求路径"
            :labelCol="labelCol"
            :wrapperCol="wrapperCol">
            <a-input v-decorator="['path', {rules: [{required: true}]}]" placeholder="请输入权限标识" />
          </a-form-item>
          <a-form-item
            label="备注信息"
            :labelCol="labelCol"
            :wrapperCol="wrapperCol" >
            <a-textarea :rows="4" v-decorator="['remarks']"></a-textarea>
          </a-form-item>
        </div>
      </a-form>
    </a-spin>
    <template slot="footer">
      <a-button key="cancel" @click="handleCancel">取消</a-button>
      <a-button key="forward" :loading="confirmLoading" type="primary" @click="handleNext()">完成</a-button>
    </template>
  </a-modal>
</template>

<script>
import pick from 'lodash.pick'
import { SaveMenuBtn, GetMenuBtn } from '@/api/sys'

let that, backFunc
export default {
  name: 'MenuBtnEdit',
  data () {
    return {
      id: null,
      menuId: '',
      projectList: [],
      projectListCache: [],
      projectTypeList: [],
      labelCol: {
        xs: { span: 24 },
        sm: { span: 7 }
      },
      wrapperCol: {
        xs: { span: 24 },
        sm: { span: 13 }
      },
      visible: false,
      confirmLoading: false,
      form: this.$form.createForm(this)
    }
  },
  methods: {
    edit (record, menuId, backFunc) {
      that = this
      that.backFunc = backFunc
      that.visible = true
      that.menuId = menuId
      // 初始化
      if (record) {
        that.id = record.id
        that.menuId = record.sys_menu_id
        GetMenuBtn({ id: record.id })
          .then(res => {
            that.loading = false
            if (res.ret === 200) {
              record = res.data
              const { form: { setFieldsValue } } = that
              that.$nextTick(() => {
                setFieldsValue(pick(record, 'name', 'permission', 'remarks', 'method', 'path'))
              })
            }
          })
      } else {
        that.id = null
        that.record = null
        that.loading = false
      }
    },
    handleNext () {
      const { form: { validateFields } } = this
      // last step
      this.confirmLoading = true
      validateFields((errors, values) => {
        console.log('errors:', errors, 'val:', values)
        if (!errors) {
          // 保存消息
          values.id = that.id
          values.sys_menu_id = that.menuId
          SaveMenuBtn(values)
            .then(res => {
              that.confirmLoading = false
              if (res.ret === 200) {
                that.$notification.success({
                  message: '系统提示',
                  description: res.msg
                })
                this.visible = false
                if (that.backFunc) {
                  that.backFunc()
                }
              }
            })
        } else {
          this.confirmLoading = false
        }
      })
    },
    handleChange (value) {
      console.log(`selected ${value}`)
    },
    handleCancel () {
      // clear form
      this.visible = false
    }
  }
}
</script>
