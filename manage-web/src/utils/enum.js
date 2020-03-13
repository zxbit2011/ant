export const UserType = { sec: 'sec', teacher: 'teacher', student: 'student' }
export const DisType = [
  { value: 'file', label: '文件' },
  { value: 'multiselect', label: '多选框' },
  { value: 'radio', label: '单选框' },
  { value: 'video', label: '视频' },
  { value: 'text', label: '文本域' },
  { value: 'score', label: '评分' },
  { value: 'tally', label: '计数' },
  { value: 'time', label: '计时' },
  { value: 'img', label: '图片' },
  { value: 'img_text', label: '图文' },
  { value: 'audio', label: '录音' },
  { value: 'title_text', label: '标题文本' }
]
/**
 * 评价方式
 * @constructor
 */
export function GetDisTypeName (type) {
  const model = DisType.filter(item => item.value === type)
  if (model.length > 0) {
    return model[0].label
  }
  return ''
}

export const ReviewStatusMap = [
  {
    value: 0,
    status: 'orange',
    text: '未审核'
  },
  {
    value: 1,
    status: 'green',
    text: '已审核'
  }
]

export function ReviewStatusText (type) {
  if (!type || type === '') return '未审核'
  const model = ReviewStatusMap.filter(item => item.value === parseInt(type))
  if (model.length > 0) {
    return model[0].text
  }
  return '未审核'
}
export function ReviewStatusColor (type) {
  if (!type || type === '') return 'orange'
  const model = ReviewStatusMap.filter(item => item.value === parseInt(type))
  if (model.length > 0) {
    return model[0].status
  }
  return 'orange'
}

export const TaskTypeMap = [
  {
    value: 1,
    status: 'orange',
    text: '单次任务'
  },
  {
    value: 2,
    status: 'green',
    text: '持续任务'
  }
]

export function TaskTypeText (type) {
  if (!type || type === '') return '单次任务'
  const model = TaskTypeMap.filter(item => item.value === parseInt(type))
  if (model.length > 0) {
    return model[0].text
  }
  return '单次任务'
}

export function TaskTypeColor (type) {
  if (!type || type === '') return 'orange'
  const model = TaskTypeMap.filter(item => item.value === parseInt(type))
  if (model.length > 0) {
    return model[0].status
  }
  return 'orange'
}
