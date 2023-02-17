let films = [{
    "Name": "Mr. and Mrs. Smith",
    "Country": "U.S.A",
    "duration": 126
},
{
    "Name":"Black Panther: Wakanda Forever",
    "Country": "U.S.A",
    "duration": 195
},
{
    "Name": "Spider-Man: No Way Home",
    "Country": "U.S.A",
    "duration": 150
}]

let sum = 0;

for(let i = 0; i < films.length; i++){
    sum += films[1]['duration']
}

avg_year = sum / films.length;

let avatar = {
    "Name": "avatar",
    "Country": "U.S.A",
    "duration": 210
}

films.push(avatar)
function write(i, films){ 
    return films[i]; 
} 
 
for (let i = 0; i < films.length; i++) { 
    let obj = write(i, films); 
    document.write("<h5>Films: " + obj["Name"] + "</h5>"); 
    document.write("<h5>Country: " + obj["Country"] + "</h5>"); 
    document.write("<h5>Duration (minutes): " + obj["duration"] + ""); 
    document.writeln("") 
} 
 
 
document.write("<h5>Average duration: " + avg(films) + " minutes</h5>")
