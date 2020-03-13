<template>
  <a-modal
    title="信息编辑"
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
            label="信息标题"
            :labelCol="labelCol"
            :wrapperCol="wrapperCol">
            <a-input v-decorator="['title', {rules: [{required: true}]}]" placeholder="请输入信息标题" />
          </a-form-item>
          <a-form-item
            label="发送范围"
            :labelCol="labelCol"
            :wrapperCol="wrapperCol" >
            <a-radio-group
              @change="onChangeType"
              v-decorator="['type', {initialValue: 1, rules: [{required: true}]}]">
              <a-radio :value="1">全部</a-radio>
            </a-radio-group>
          </a-form-item>
          <a-form-item
            label="信息内容"
            :labelCol="labelCol"
            :wrapperCol="wrapperCol" >
            <a-textarea :rows="4" v-decorator="['content', {rules: [{required: true}]}]"></a-textarea>
          </a-form-item>
          <a-form-item
            :labelCol="labelCol"
            :wrapperCol="wrapperCol">
            <span slot="label">
              信息附件&nbsp;
              <a-tooltip>
                <template slot="title">
                  <div><b>文件格式限制：</b></div>
                  <div>图片：jpg,png,gif,tif,bmp,psd,ico；</div>
                  <div>音频：mp3,wav,amr,m4a,ogg,flac；</div>
                  <div>视频：mp4,m4v,mkv,mov,avi,flv；</div>
                  <div>文档：doc,docx,xls,xlsx,ppt,pptx；</div>
                </template>
                <a-icon type="question-circle-o"/>
              </a-tooltip>
            </span>
            <a-upload
              action="/api/auth/files/upload/auth?dir=notify"
              listType="picture"
              :fileList="fileList"
              class="upload-list-inline"
              :headers="headers"
              :beforeUpload="handleFileBeforeUpload"
              @change="handleFileChange"
              @preview="previewFile"
              :remove="handleFileRemove">
              <a-button>
                <a-icon type="upload"/>
                点击上传
              </a-button>
            </a-upload>
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
import Vue from 'vue'
import { ACCESS_TOKEN } from '@/store/mutation-types'
import { DelSysNotifyFile, GetSysNotifyInfo, SaveSysNotify } from '@/api/notify'
import { removeFilesLog } from '@/api/filelog'
let that, backFunc
export default {
  name: 'QuotaEdit',
  data () {
    return {
      id: null,
      years: [],
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
      form: this.$form.createForm(this),
      typeOptions: [
        { label: '全部', value: '1' },
        { label: '指定项目', value: '2' }
      ],
      typeCheckedValue: 0,
      fileList: [],
      headers: {
        'X-Requested-With': 'XMLHttpRequest',
        'Access-Token': Vue.ls.get(ACCESS_TOKEN)
      }
    }
  },
  methods: {
    edit (record, backFunc) {
      that = this
      this.backFunc = backFunc
      this.visible = true
      // 初始化
      this.typeCheckedValue = 0
      if (record) {
        that.id = record.id
        GetSysNotifyInfo({ id: record.id })
          .then(res => {
            that.loading = false
            if (res.ret === 200) {
              record = res.data.notify
              record.type = parseInt(record.type)
              that.typeCheckedValue = record.type
              that.Form = res.data.notify
              const { form: { setFieldsValue } } = that
              if (res.data.notifyFile.length > 0) {
                this.fileList = res.data.notifyFile.map(option => {
                  return {
                    uid: option.id,
                    id: option.id,
                    name: option.name,
                    status: 'done',
                    url: option.url,
                    thumbUrl: option.url
                  }
                })
              }
              that.$nextTick(() => {
                setFieldsValue(pick(record, 'title', 'type', 'content', 'obj_id', 'status'))
              })
            }
          })
      } else {
        that.id = null
        that.record = null
        that.fileList = null
        that.loading = false
      }
    },
    previewFile (info) {
      console.log(info)
      window.open(info.url, '_blank')
    },
    handleFileChange (info) {
      let fileList = [...info.fileList]
      fileList = fileList.map((file) => {
        if (file.response) {
          if (file.response.ret !== 200) {
            that.$message.error(file.response.msg)
            return null
          }
          // Component will show file.url as link
          file.url = file.response.data.url
          file.uid = file.response.data.id
          file.id = -1
        }
        return file
      })
      that.fileList = fileList.filter(item => item !== null)
      if (info.file.status === 'done' && info.response && info.response.ret === 200) {
        that.$message.success(`${info.file.name} 上传文件成功`)
      } else if (info.file.status === 'error') {
        that.$message.error(`${info.file.name} 上传文件失败`)
      }
    },
    handleFileBeforeUpload (file) {
      if (file.size === 0) {
        that.$message.warning('空文件，请重新上传', 30)
        return false
      }
      const isLt2M = file.size / 1024 / 1024 < 30
      if (!isLt2M) {
        that.$message.warning('文件大小不能超过 30MB!', 30)
      }
      return isLt2M
    },
    handleFileRemove (file) {
      this.$confirm({
        title: '任务附件删除',
        content: `是否确定删除此任务附件？`,
        okText: '确认',
        cancelText: '取消',
        onOk () {
          if (file.id && file.id !== -1) {
            // 真实的上传文件，删除提示
            DelSysNotifyFile({ id: file.id })
              .then(res => {
                if (res.ret === 200) {
                  that.$message.success(`删除文件成功`)
                  that.fileList = that.fileList.filter(item => item.uid !== file.uid)
                }
              })
          } else {
            // 临时上传
            removeFilesLog({ id: file.uid })
              .then(res => {
                if (res.ret === 200) {
                  that.$message.success(`删除文件成功`)
                  that.fileList = that.fileList.filter(item => item.uid !== file.uid)
                }
              })
          }
        }
      })
      return false
    },
    onChangeType (checkedValue) {
      console.log('checked = ', checkedValue)
      console.log('value = ', checkedValue.target.value)
      this.typeCheckedValue = checkedValue.target.value
    },
    filterOption (input, option) {
      return option.componentOptions.children[0].text.toUpperCase().indexOf(input.toUpperCase()) >= 0
    },
    filter (inputValue, path) {
      return (path.some(option => (option.label).toLowerCase().indexOf(inputValue.toLowerCase()) > -1))
    },
    handleNext () {
      const { form: { validateFields } } = this
      // last step
      this.confirmLoading = true
      validateFields((errors, values) => {
        console.log('errors:', errors, 'val:', values)
        if (!errors) {
          // 保存消息
          const fileList = that.fileList === null ? [] : that.fileList.filter(item => item.status === 'done' & item.id === -1).map(option => {
            return {
              uid: option.uid,
              name: option.name,
              status: option.status,
              url: option.url
            }
          })
          values.id = that.id
          values.fileList = fileList.length === 0 ? '' : JSON.stringify(fileList)
          SaveSysNotify(values)
            .then(res => {
              that.confirmLoading = false
              if (res.ret === 200) {
                that.id = res.data.id
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
