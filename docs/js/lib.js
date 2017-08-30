
var jclubsheet = 'https://docs.google.com/spreadsheets/d/1M38X99dQ_rLCaE6mqX9ct2DF-YO6ySAfd3f1_xwerYs/edit#gid=0';
var jclubtakeaways = 'https://docs.google.com/spreadsheets/d/1MpfQ3BaInYKHI6ebIJEYml1xgfAM0rr5ymmDi6xs8aw/edit#gid=0';

function get_date_href(date) {
    return "#" + date.replace(/\//g,"-");
}
function parse_date_href(date) {
    return date.split('#')[1].replace(/-/g,"/");
}

function get_row_from_date(date, callback) {
   sheetrock({
      url: jclubsheet,
      query: 'select A,B,C',
      callback: function (error, options, response) {
        for (i in response.rows) {
            var row = response.rows[i];
            var labels = row.labels;
            var cells = row.cellsArray;
            var rowobj = {};
            for (j in labels) {
                rowobj[labels[j]] = cells[j];
            }
            if (rowobj["Date"] == date) {
                callback(rowobj)
                return;
            }
        }
        callback(null);
      }
   });
}
function format_date(date) {
    return (date.getMonth() + 1) + "/" + date.getDate() + "/" + (date.getYear() + 1900);
}
function get_last_thursday() {
    var now = new Date();
    var day = now.getDay();

    var daysAfterLastThursday;
    if (day < 4)
        daysAfterLastThursday = -3 - day;
    else if (day > 4)
        daysAfterLastThursday = 4 - day;
    else 
        daysAfterLastThursday = 0;

    var currentMs = now.getTime();
    var lastThursday = new Date(currentMs + (daysAfterLastThursday * 24 * 60 * 60 * 1000));
    return format_date(lastThursday);
}
