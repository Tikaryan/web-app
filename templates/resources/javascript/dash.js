var name, userId = ""
$(document).ready(function(){
      let cookie =  document.cookie.split("|")
      for (let i=0; i < cookie.length; i++){
             let cook = cookie[i].split("=")
             for (let j = 0; j < cook[j]; j++){
                    if(/name/.test(cook[j]))
                      name =  cook[j+1]
                      else if(/userid/.test(cook[j]))
                            userId = cook[j+1]
             }
      }

      $('#name').text(name)
      $('#userId').text(userId)
})