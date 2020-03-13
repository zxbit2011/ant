export function timeFix () {
  const time = new Date()
  const hour = time.getHours()
  return hour < 9 ? '早上好' : hour <= 11 ? '上午好' : hour <= 13 ? '中午好' : hour < 20 ? '下午好' : '晚上好'
}

export function welcome () {
  const arr = ['欢迎回来！']
  const index = Math.floor(Math.random() * arr.length)
  return arr[index]
}

/**
 * 触发 window.resize
 */
export function triggerWindowResizeEvent () {
  const event = document.createEvent('HTMLEvents')
  event.initEvent('resize', true, true)
  event.eventType = 'message'
  window.dispatchEvent(event)
}

export function handleScrollHeader (callback) {
  let timer = 0

  let beforeScrollTop = window.pageYOffset
  callback = callback || function () {}
  window.addEventListener(
    'scroll',
    event => {
      clearTimeout(timer)
      timer = setTimeout(() => {
        let direction = 'up'
        const afterScrollTop = window.pageYOffset
        const delta = afterScrollTop - beforeScrollTop
        if (delta === 0) {
          return false
        }
        direction = delta > 0 ? 'down' : 'up'
        callback(direction)
        beforeScrollTop = afterScrollTop
      }, 50)
    },
    false
  )
}

/**
 * Remove loading animate
 * @param id parent element id or class
 * @param timeout
 */
export function removeLoadingAnimate (id = '', timeout = 1500) {
  if (id === '') {
    return
  }
  setTimeout(() => {
    document.body.removeChild(document.getElementById(id))
  }, timeout)
}

/**
 格式化时间
 */
export function formatDate (date, fmt) {
  if (date === '' || date === null || date === undefined) {
    return null
  }
  if (fmt === '' || fmt == null || fmt === undefined) {
    fmt = 'yyyy-MM-dd hh:mm:ss'
  }
  date = new Date(date)
  var o = {
    'M+': date.getMonth() + 1, // 月份
    'd+': date.getDate(), // 日
    'h+': date.getHours(), // 小时
    'm+': date.getMinutes(), // 分
    's+': date.getSeconds(), // 秒
    'q+': Math.floor((date.getMonth() + 3) / 3), // 季度
    'S': date.getMilliseconds() // 毫秒
  }
  if (/(y+)/.test(fmt)) fmt = fmt.replace(RegExp.$1, (date.getFullYear() + '').substr(4 - RegExp.$1.length))
  for (var k in o) { if (new RegExp('(' + k + ')').test(fmt)) fmt = fmt.replace(RegExp.$1, (RegExp.$1.length === 1) ? (o[k]) : (('00' + o[k]).substr(('' + o[k]).length))) }
  return fmt
}

const fileExt = [
  '.doc',
  '.docx',
  '.xls',
  '.xlsx',
  '.txt',
  '.pdf',
  '.hlp',
  '.wps',
  '.ppt',
  '.ppt',
  '.pps'
]

const imgExt = [
  '.png',
  '.jpeg',
  '.bmp',
  '.jpg',
  '.gif'
]

const videoExt = [
  '.avi',
  '.rm',
  '.rmvb',
  '.mpeg',
  '.mpg',
  '.ogg',
  '.ogv',
  '.mov',
  '.wmv',
  '.webm'
]
const audioExt = [
  '.amr',
  '.mp3',
  '.mp4',
  '.wav'
]
// 获取文件名后缀名
export function extension (str) {
  let ext = ''
  const name = str.toLowerCase()
  const i = name.lastIndexOf('.')
  if (i > -1) {
    ext = name.substring(i)
  }
  if (ext.indexOf('?') > -1) {
    ext = ext.substr(0, ext.indexOf('?'))
  }
  return ext
}

// 判断Array中是否包含某个值
export function contain (str, obj) {
  for (let i = 0; i < str.length; i++) {
    if (str[i] === obj) { return true }
  }
  return false
};

export function extMatch (str, extType) {
  if (contain(extType, extension(str))) {
    return true
  } else {
    return false
  }
}

export function extType (str) {
  if (contain(fileExt, extension(str))) {
    return 'file'
  } else if (contain(imgExt, extension(str))) {
    return 'img'
  } else if (contain(videoExt, extension(str))) {
    return 'video'
  } else if (contain(audioExt, extension(str))) {
    return 'audio'
  } else {
    return 'default'
  }
}

export function convertTime (time) {
  if (time !== null && time !== '') {
    if (time > 60 && time < 60 * 60) {
      time = parseInt(time / 60.0) + '分钟' + parseInt((parseFloat(time / 60.0) -
        parseInt(time / 60.0)) * 60) + '秒'
    } else if (time >= 60 * 60 && time < 60 * 60 * 24) {
      time = parseInt(time / 3600.0) + '小时' + parseInt((parseFloat(time / 3600.0) -
        parseInt(time / 3600.0)) * 60) + '分钟' +
        parseInt((parseFloat((parseFloat(time / 3600.0) - parseInt(time / 3600.0)) * 60) -
          parseInt((parseFloat(time / 3600.0) - parseInt(time / 3600.0)) * 60)) * 60) + '秒'
    } else if (time >= 60 * 60 * 24) {
      time = parseInt(time / 3600.0 / 24) + '天' + parseInt((parseFloat(time / 3600.0 / 24) -
        parseInt(time / 3600.0 / 24)) * 24) + '小时' + parseInt((parseFloat(time / 3600.0) -
        parseInt(time / 3600.0)) * 60) + '分钟' +
        parseInt((parseFloat((parseFloat(time / 3600.0) - parseInt(time / 3600.0)) * 60) -
          parseInt((parseFloat(time / 3600.0) - parseInt(time / 3600.0)) * 60)) * 60) + '秒'
    } else {
      time = parseInt(time) + '秒'
    }
  } else {
    return '0'
  }
  return time
}
