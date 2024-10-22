class PostController < ApplicationController
  def index
    posts = load_markdown_posts
    @recent_posts = posts.sort_by { |post| post[:date] }.reverse.first(5)

    @metadata = {
      "title" => "Últimos Posts",
      "description" => "Confira os meus últimos posts e atualizações no blog."
    }
  end

  def show
    file_name = params[:file_name]
    file_path = Rails.root.join("app", "markdown_files", "#{file_name}.md")

    if File.exist?(file_path)
      content = File.read(file_path)

      if content =~ /\A---(.|\n)*?---/
        front_matter = content.match(/\A---(.|\n)*?---/)[0]
        @metadata = YAML.safe_load(front_matter)
        @markdown_content = content.sub(front_matter, "").strip
      else
        @metadata = {}
        @markdown_content = content
      end
    else
      render plain: "File not found", status: :not_found
    end
  end

  private

  def load_markdown_posts
    posts_directory = Rails.root.join("app", "markdown_files")
    Dir.glob("#{posts_directory}/*.md").map do |file|
      content = File.read(file)

      # Extrai os metadados do front matter
      if content =~ /\A---(.|\n)*?---/
        front_matter = content.match(/\A---(.|\n)*?---/)[0]
        metadata = YAML.safe_load(front_matter)

        # Extrai a data e o título do nome do arquivo
        file_name = File.basename(file, ".md")
        date_string = file_name.split("-").first(3).join("-") # Extrai a data
        date = Date.parse(date_string) # Converte para um objeto Date

        # O título é a parte do nome do arquivo após a data
        title = file_name.split("-")[3..].join(" ").capitalize

        {
          title: title,
          description: metadata["description"],
          author: metadata["author"],
          date: date,
          file_name: file_name
        }
      end
    end.compact
  end
end
