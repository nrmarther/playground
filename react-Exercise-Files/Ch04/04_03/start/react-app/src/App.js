import "./App.css";

const [firstCity, second] = ["Tokyo", "Tahoe City", "Bend"];

console.log(firstCity);
console.log(second);

function App({ firstCity }) {
  return (
    <div className="App">
      <h1>Hello from {firstCity}</h1>
    </div>
  );
}

export default App;
