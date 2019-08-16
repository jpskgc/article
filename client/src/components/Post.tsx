import React from 'react';
import {Form, Container, List, ListItemProps, Confirm} from 'semantic-ui-react';
import axios from 'axios';
import {Redirect} from 'react-router';
import Dropzone from 'react-dropzone';

interface ArticleState {
  title: string;
  content: string;
  redirect: boolean;
  files: File[];
  confirm: boolean;
}

class Post extends React.Component<{}, ArticleState> {
  constructor(props: {}) {
    super(props);
    this.state = {
      title: '',
      content: '',
      redirect: false,
      files: [],
      confirm: false,
    };
    this.handleChangeTitle = this.handleChangeTitle.bind(this);
    this.handleChangeContent = this.handleChangeContent.bind(this);
    this.handleSubmit = this.handleSubmit.bind(this);
    this.renderRedirect = this.renderRedirect.bind(this);
    this.handleOnDrop = this.handleOnDrop.bind(this);
    this.handleRemove = this.handleRemove.bind(this);
    this.handleConfirmOpen = this.handleConfirmOpen.bind(this);
    this.handleConfirmClose = this.handleConfirmClose.bind(this);
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

  handleRemove(
    event: React.MouseEvent<HTMLAnchorElement, MouseEvent>,
    data: ListItemProps
  ) {
    const fileName = data.content;
    const targetFile = this.state.files.find(v => v.name === fileName);
    const index = this.state.files.indexOf(targetFile as File);
    this.state.files.splice(index, 1);
    this.setState({files: this.state.files});
  }

  handleConfirmOpen() {
    this.setState({confirm: true});
  }

  handleConfirmClose() {
    this.setState({confirm: false});
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
          {/* Fix desigin*/}
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
          <List>
            {(this.state.files || []).map((file, i) => {
              return (
                <List.Item
                  icon="image"
                  content={file.name}
                  onClick={this.handleRemove}
                />
              );
            })}
          </List>
          <Confirm
            open={this.state.confirm}
            onCancel={this.handleConfirmClose}
            onConfirm={this.handleConfirmOpen}
          />
          <Form.Button content="Submit" onClick={this.handleSubmit} />
        </Form>
      </Container>
    );
  }
}

export default Post;
