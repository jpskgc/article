import React from 'react';
import {Form, Container, List, ListItemProps, Icon} from 'semantic-ui-react';
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
    this.handleRemove = this.handleRemove.bind(this);
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
    //const fileName = data.content;
    console.log(data.value);
    const fileName = data.value;
    const targetFile = this.state.files.find(v => v.name === fileName);
    const index = this.state.files.indexOf(targetFile as File);
    this.state.files.splice(index, 1);
    this.setState({files: this.state.files});
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

  componentDidMount() {
    this.setState(() => window.scrollTo(0, 0));
  }

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
          <Dropzone accept="image/*" onDrop={this.handleOnDrop}>
            {({getRootProps, getInputProps, open}) => (
              <section>
                <div
                  {...getRootProps()}
                  style={{
                    margin: '15px auto',
                    flex: 1,
                    display: 'flex',
                    flexDirection: 'column',
                    alignItems: 'center',
                    padding: '20px',
                    borderWidth: 2,
                    borderRadius: 2,
                    borderColor: '#eeeeee',
                    borderStyle: 'dashed',
                    backgroundColor: '#fafafa',
                    color: '#bdbdbd',
                    outline: 'none',
                    transition: 'border .24s ease-in-out',
                  }}
                >
                  <input {...getInputProps()} />
                  <p>
                    <Icon name="image" />
                    Drag 'n' drop some files here, or click to select files
                  </p>
                  <button type="button" onClick={open}>
                    Open File Dialog
                  </button>
                </div>
              </section>
            )}
          </Dropzone>
          {(this.state.files || []).map((file, i) => {
            return (
              <List horizontal>
                <List.Item icon="image" content={file.name} />
                <List.Item
                  icon="delete"
                  content="cancel"
                  value={file.name}
                  onClick={this.handleRemove}
                />
              </List>
            );
          })}
          <Form.Button content="Submit" onClick={this.handleSubmit} />
        </Form>
      </Container>
    );
  }
}

export default Post;
