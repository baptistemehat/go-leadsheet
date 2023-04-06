import React from 'react';
import ReactDOM from 'react-dom/client';
import './index.css';
import App from './App';
import reportWebVitals from './reportWebVitals';

class SongForm extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      metadata: [],
      leadsheet: 'Type your lyrics and chords here',
    };

    this.handleLeadsheetChange = this.handleLeadsheetChange.bind(this);
    this.handleSubmit = this.handleSubmit.bind(this);
  }

  handleLeadsheetChange(event) {
    this.setState({leadsheet: event.target.value});
  }

  handleSubmit(event) {
    console.log(this.state.leadsheet);
    event.preventDefault();
  }

  render() {
    return(
      <form onSubmit={this.handleSubmit}>
        <label>Title:
          <textarea value={this.state.leadsheet} onChange={this.handleLeadsheetChange} />
        </label>
        <input type="submit" value="Submit" on />
      </form>
    );
  }

}

const root = ReactDOM.createRoot(document.getElementById('root'));
root.render(
  <React.StrictMode>
    <App />
  </React.StrictMode>
);

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
reportWebVitals();
