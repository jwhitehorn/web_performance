class RestApi < Sinatra::Base
  register Sinatra::Async
  before { content_type 'application/json' }

  aget '/sets/:name/cards' do
    set = URI.decode(params[:name])
    DataAccess.begin do |dao|
      dao.cards_in_set(set) do |cards|
        body cards
        dao.close
      end
    end
  end

end


map '/api' do
  run RestApi.new
end
