class DataAccess
  
  def self.begin(&block)
    block.call DataAccess.new
  end
  
  def cards_in_set(name, &block)
    set_name = escape(name)
    redis.smembers("set-cards-#{set_name}") do |keys|
      card_keys = keys.map { |key| "card-#{key}" }
      objects_for_keys(card_keys) { |objs| block.call objs }
    end
  end
  
  def close
    redis.close_connection
  end
  
private  
  def redis
    @redis ||= EM::Hiredis.connect
  end
  
  def escape(string)
    string.gsub(/\W/, "-").downcase.gsub(/-+/, '-')
  end
  
  def objects_for_keys(keys, &block)
    if keys.length > 0
      redis.mget(*keys) do |objs|
        #multiple returned objects need to be added to JSON array
        block.call "[#{objs.join(", ")}]"
      end
    else
      block.call "[]"
    end
  end
    
  
end