class PagesController < ApplicationController
  def about
    @metadata = {
      "title" => "Sobre Mim"
    }
  end
end
