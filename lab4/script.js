let dating = [
  { "place": "Oi-Qaragai", "date": "31.08.2021", "money": 20000 },
  { "place": "Kolsai-Kayyndy", "date": "01.06.2022", "money": 30000 },
  { "place": "Medey", "date": "14.02.2023", "money": 15000 }
];

let sum = dating.reduce((accumulator, currentValue) => {
  return accumulator + currentValue.money;
}, 0);

console.log(sum);


function createDating(place, date, money) {
  let dating = {
    place: place,
    date: date,
    money: money
  };

  return dating;
}

let place = createDating("Oi-Qaragai", 31, "20000");

console.log(place.place);
console.log(place.date);
console.log(place.money);


document.write("<h2>Our best dates</h2>");

dating.forEach((date) => {
  document.write("<p>" + date.place + "<br>" + date.date + "<br>" + date.money + "<br>" + "</p>");
});


