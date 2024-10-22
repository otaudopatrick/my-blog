class PostController < ApplicationController
  def index
    @recent_posts = Post.all.order(created_at: :desc).limit(5)
  end

  def show
    @post = Post.find_by(slug: params[:slug])

    if @post
      @markdown_content = @post.body
      @metadata = {
        "title" => @post.title,
        "description" => @post.description,
        "author" => @post.author
      }
    else
      render plain: "Post não encontrado", status: :not_found
    end
  end
end
