import React from 'react';

/**
 * Contains the components required to input a song
 */
class SongForm extends React.Component {
    constructor(props) {
      super(props);
      this.state = {
  
        /**
         * Song metadata (title, composer, etc.)
         */
        metadata: {
          title: '',
          composer: '',
        },
  
        /**
         * Content of the 'leadsheet' text area
         */
        leadsheet: 'Type your lyrics and chords here',

        /**
         * Displayed generation status
         */
        generationStatus: 'not started'
      };
  
      this.handleLeadsheetChange = this.handleLeadsheetChange.bind(this);
      this.handleMetadataChange = this.handleMetadataChange.bind(this)
      this.handleSubmit = this.handleSubmit.bind(this);
      this.handleCheckGenerationStatus = this.handleCheckGenerationStatus.bind(this)
    }
  
    /**
     * Handles changes in the leadsheet text area
     * @param {Event} event 
     */
    handleLeadsheetChange(event) {
      this.setState({leadsheet: event.target.value});
    }
  
    /**
     * Handles changes in metadata text fields (title, composer, etc.)
     * @param {Event} event 
     */
    handleMetadataChange(event) {
  
      // create copy of current metadata and merge the new value into it
      var newMetadata = {
        ...this.state.metadata,
        [event.target.name] : event.target.value
      }
  
      // update current metadata
      this.setState({metadata: newMetadata});
    }
  
    // TODO : implement
    /**
     * Checks if input data is valid
     * This method would eventually check for syntax error in leadsheet text area
     * @returns true for now
     */
    checkDataValidity() {
      return true;
    }
  
  
    /**
     * Handles submission
     * @param {Event} event 
     */
    handleSubmit(event) {
      event.preventDefault();
      
      // if user data is valid
      if (this.checkDataValidity()) {
        // send to core api
        fetch('http://localhost:8000/api/song?type=text', 
          {
            method : 'POST',
            body: JSON.stringify({
              leadsheet: this.state.leadsheet
            })                    
          })
        .then(response => response.text())
        .then(text => console.log(text))
        .catch(err => {
          console.log(err);
        });
      }
      // TODO : if user data is not valid
    }
  
    // TODO : make this method a 'deamon' that pulls generation status periodically
    /**
     * Handles press on 'get generation status' button
     * This method would eventually pull generation status periodically to update UI
     * and 'unlock' the download button once generation is complete 
     * @param {Event} event 
     */
    handleCheckGenerationStatus(event) {
      fetch('http://localhost:8000/api/status', {method : 'GET'})
        .then(response => response.text())
        .then(text => {
          console.log(text)
          this.setState({ generationStatus: text })
        })
        .catch(err => {
          console.log(err);
        });
    }
  
    /**
     * Handle pdf file download
     * @param {Event} event 
     */
    handleDownloadPdf(event) {
      fetch('http://localhost:8000/api/song', {method : 'GET'})
      .then(response => {
        if (response.status < 300)
          return response.blob()
        // TODO : do this better
        throw new Error("no file yet")
      })
      .then(blob => {
        let url = window.URL.createObjectURL(blob);
        let a = document.createElement('a');
        a.href = url;
        a.download = 'test.pdf';
        a.click();
      })
      .catch(err => {
        console.log(err);
      });
    }
    


    render() {
      return(
        
        <form className="SongForm" onSubmit={this.handleSubmit}>
          <div className='Metadata-fields'>
            <div>
                <label>Title:
                <input 
                    name="title"
                    type="text" 
                    value={this.state.metadata.title} 
                    onChange={this.handleMetadataChange} />
                </label>
            </div>
            
            <div>
                <label>Composer:
                <input
                    name="composer"
                    type="text"
                    value={this.state.metadata.composer}
                    onChange={this.handleMetadataChange} />
                </label>
            </div>
          </div>
           
  
          <div className='Leadsheet-fields'>
            <div>
                <label>Leadsheet:</label>
            </div>
            <div>
                <textarea className='leadsheet-text-area' value={this.state.leadsheet} onChange={this.handleLeadsheetChange} />
            </div>
          </div>
  
          <input type="submit" value="Generate" />
        
          <div>
            <input className="" type="button" value="Check generation status" onClick={this.handleCheckGenerationStatus} />
            <label>Generation {this.state.generationStatus}</label>
          </div>
         
         <div>
            <input type="button" value="Download file" onClick={this.handleDownloadPdf} />
         </div>
        </form>
      );
    }
  }

export default SongForm;