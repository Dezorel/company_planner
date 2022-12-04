/* When the user clicks on the button,
toggle between hiding and showing the dropdown content */
function myFunction() {
  document.getElementById("myDropdown").classList.toggle("show");

  let nameCompany = document.getElementById("company_name").value

  var data = []
      newURL = requestURL+'/cabinet/' + nameCompany
      async function request(){
         data = await sendRequest('GET', newURL)
      }
      request()

      setTimeout(()=>{
         console.log(data)
         let cabList = document.getElementById("myDropdown")

         data.forEach(cab => {
          let element = document.createElement("a")
          element.append(cab)

        })

      },1000)


}

function onloadCalendar(companyName){

  var data = []
      newURL = requestURL+'/schedule/'+companyName
      async function request(){
         data = await sendRequest('GET', newURL) 
      }
      request()

      setTimeout(()=>{
         console.log(data)

          var calendarEl = document.getElementById('calendar');

          var calendar = new FullCalendar.Calendar(calendarEl, {
              headerToolbar: {
                  left: 'prevYear,prev,next,nextYear today',
                  center: 'title',
                  right: 'dayGridMonth,dayGridWeek,dayGridDay'
              },
              initialDate: '2020-09-12',
              navLinks: true, // can click day/week names to navigate views
              editable: true,
              dayMaxEvents: true, // allow "more" link when too many events
              events: [
                  {
                      "title":data[0].cabinet_number,
                      "start":data[0].date_time_start,
                      "end":data[0].date_time_end,
                  }
              ]
          });

          calendar.render()

      },1000)
      
}

function filterFunction() {
  var input, filter, ul, li, a, i;
  input = document.getElementById("myInput");
  filter = input.value.toUpperCase();
  div = document.getElementById("myDropdown");
  a = div.getElementsByTagName("a");
  for (i = 0; i < a.length; i++) {
    txtValue = a[i].textContent || a[i].innerText;
    if (txtValue.toUpperCase().indexOf(filter) > -1) {
      a[i].style.display = "";
    } else {
      a[i].style.display = "none";
    }
  }
}
// показать скрыть кабинет
function showCabinet() {
  document.getElementById("calendar__container").setAttribute("style", "display: none");
  document.getElementById("cabinet").setAttribute("style", "display: block");
}
// показать скрыть календарь
function showCalendar() {
  document.getElementById("cabinet").setAttribute("style", "display: none");
  document.getElementById("calendar__container").setAttribute("style", "display: block");
}

// function showPage2() {
//   document.getElementById("page1").setAttribute("style", "display: none");
//   document.getElementById("page2").setAttribute("style", "display: block");
// }

// Show cabinet info
function createCabinet() {
  let nameCompany = document.getElementById("cabCompanyName").value
  let cabNum = document.getElementById('cabNum').value
  let size = document.getElementById('cabSize').value
  let optionWiFI = document.getElementById('wifiChekBox').checked
  let optionProector = document.getElementById('proiectorChekBox').checked
  let property = ""
    
  if (optionWiFI === true) {
      property+= " Wifi "
  }
  if (optionProector === true) {
      property+= " Proector "
  }

  var data = []
      newURL = requestURL+'/cabinet'
      async function request(){
         data = await sendRequest('POST', newURL, {number:cabNum, company:nameCompany, size:size, property:property })
      }
      request()

      setTimeout(()=>{
         console.log(data)
      },1000)
      
  return console.log(cabNum)
}

function showPage2(){
  let nameCompany = document.getElementById("company_name").value
  console.log(nameCompany)
  var data = []
      newURL = requestURL+'/company/'+nameCompany
      async function request(){
         data = await sendRequest('GET', newURL)
      }
      request()

      setTimeout(()=>{
         console.log(data)
          onloadCalendar(data.name)
      },1000)
}
function createCompany(){
  let nameCompany = document.getElementById("company_name").value
  var data = []
      newURL = requestURL+'/company'
      async function request(){
         data = await sendRequest('POST', newURL, {name:nameCompany}) 
      }
      request()

      setTimeout(()=>{
         console.log(data)
          onloadCalendar(data.name)
      },1000)
}
function addCalendarEvent(){
  let startDate = document.getElementById("startDate").value
  let endDate = document.getElementById("endDate").value
  let nameCompany = document.getElementById("companyName").value
  let cabNumber = document.getElementById("cabNumber").value
  var data = []
      newURL = requestURL+'/schedule'
      async function request(){
         data = await sendRequest('POST', newURL, {cabinet_number:cabNumber,company_name:nameCompany,date_time_start:startDate,date_time_end:endDate})
      }
      request()

      setTimeout(()=>{
         console.log(data)
      },1000)
}



