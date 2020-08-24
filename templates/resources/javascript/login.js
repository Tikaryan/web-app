$(document).ready(function () {
       var cook = document.cookie.split('|');
       console.log(cook)
       if(cook[0] != "" ){

              let data = cook[1].split('=')
              let msg = '', errMsg = ''
              if (/msg/.test(data[0])) {
                     msg = data[1]
                     $('#alert-suc').find('strong').after(msg)
                     $('#alert-suc').show();
       
              }
              else if (/error/.test(data[0])){
                     errMsg = data[1]
                     $('#alert-dang').find('strong').after(errMsg)
                     $('#alert-dang').show();
              }
       }
       document.cookie = "temp-cookie=; expires=Thu, 01 Jan 1970 00:00:00 UTC; path=/;";
       document.cookie = "session=; expires=Thu, 01 Jan 1970 00:00:00 UTC; path=/;";
})

$.fn.customSubmit = function(e){
       e.preventDefault()
       var flag = true
       $('form :input').each(function(e){
              if($(this).hasClass('is-invalid')){
                     flag = false
                     return false
              }
       })
       if(flag != false){
            let loginid =   $('#login').val().replace(/\s/g,'')
            let pass = $('#password').val().replace(/\s/g,'')
              $.ajax({
                     url: "loginAuth", type:"POST",data: {"loginid": loginid, "password": pass},dataType: "JSON",
                     success: function(res){
                            var maps = res.Data;
                            if(maps.hasOwnProperty("error")){
                                   if( $('#alert-suc').is(":visible"))
                                          $('#alert-suc').hide()
                                   $('#alert-dang').find('strong').after("Credentials are wrong");
                                   $('#alert-dang').show();
                                   $('form').each(function(){
                                          $(this).find(':input').addClass('is-invalid');
                                   });
                            }
                            else if(maps.hasOwnProperty("success"))
                                   $('#form').submit()
                     }
              })
       }
}
$(document).on('change','.is-invalid',function(){
      $('.is-invalid').each(function(){
             $(this).removeClass('is-invalid')
      })
       $('.alert').hide();
       $('#alert-dang').contents().filter(function(){
              return this.nodeType === Node.TEXT_NODE
       }).remove();
})