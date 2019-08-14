import React from 'react';
import {Form, Container} from 'semantic-ui-react';
import axios from 'axios';
import {Redirect} from 'react-router';
import Dropzone from 'react-dropzone';

interface ArticleState {
  title: string;
  content: string;
  redirect: boolean;
  files: File[];
}

class Post extends React.Component<{}, ArticleState> {
  constructor(props: {}) {
    super(props);
    this.state = {
      title: '',
      content: '',
      redirect: false,
      files: [],
    };
    this.handleChangeTitle = this.handleChangeTitle.bind(this);
    this.handleChangeContent = this.handleChangeContent.bind(this);
    this.handleSubmit = this.handleSubmit.bind(this);
    this.renderRedirect = this.renderRedirect.bind(this);
    this.handleOnDrop = this.handleOnDrop.bind(this);
  }

  handleOnDrop(acceptedFiles: File[]) {
    this.setState({files: this.state.files.concat(acceptedFiles)});
  }

  handleChangeTitle(e: React.FormEvent<HTMLInputElement>) {
    this.setState({title: e.currentTarget.value});
  }

  handleChangeContent(e: React.FormEvent<HTMLInputElement>) {
    this.setState({content: e.currentTarget.value});
  }

  async handleSubmit() {
    this.setState({
      redirect: true,
    });

    const data = {
      title: this.state.title,
      content: this.state.content,
    };

    const res = await axios.post('/api/post', data);

    const formData = new FormData();
    for (var i in this.state.files) {
      formData.append('images[]', this.state.files[i]);
    }

    const resImageNames = await axios.post('/api/post/image', formData, {
      headers: {'Content-Type': 'multipart/form-data'},
    });

    const imageData = {
      articleUUID: res.data.uuid,
      imageNames: resImageNames.data,
    };

    axios.post('/api/post/image/db', imageData).then(res => {
      console.log(res);
    });
  }

  renderRedirect = () => {
    if (this.state.redirect) {
      return <Redirect to="/post/finish" />;
    }
  };

  render() {
    return (
      <Container text style={{marginTop: '3em'}}>
        <Form>
          <Form.Input
            label="Title"
            placeholder=""
            name="title"
            value={this.state.title}
            onChange={this.handleChangeTitle}
          />
          <Form.Field
            label="Content"
            placeholder=""
            name="content"
            value={this.state.content}
            rows="20"
            control="textarea"
            onChange={this.handleChangeContent}
          />
          {this.renderRedirect()}
          {/* <Form.Button icon="upload" content="Upload Image" floated="right" />
          <input type="file" id="file" hidden /> */}
          <Dropzone accept="image/*" onDrop={this.handleOnDrop}>
            {({getRootProps, getInputProps, open}) => (
              <section>
                <div {...getRootProps()} style={{margin: '20px auto'}}>
                  <input {...getInputProps()} />
                  <p>Drag 'n' drop some files here, or click to select files</p>
                  <button type="button" onClick={open}>
                    Open File Dialog
                  </button>
                </div>
              </section>
            )}
          </Dropzone>
          <Form.Button content="Submit" onClick={this.handleSubmit} />
        </Form>
      </Container>
    );
  }
}

export default Post;
