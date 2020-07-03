import React, {Component} from 'react';
import './App.css';
import axios from "axios"

class App extends Component {
  state = {}
  
  componentDidMount() {
    axios.get("/user")
      .then(
        res => {
          console.log(res.data);
          let items = res.data.map(usr => {
            let source = `data:image/png;base64,${usr.indenticon}`
            return <li key={usr.email} className="listItem">{usr.email}: <div/> <img src={source} alt="what" /></li>
          })
          let images = <ul className="App">{items}</ul>
          this.setState({
            images: images
          })
        }
      );
  }

  render() {    
    return (
      <div className="App">
        {this.state.images}
      </div>
    )
  };
}

export default App;
