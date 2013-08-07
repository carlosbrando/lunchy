#!/usr/bin/env ruby -Ku
# require ENV['TM_SUPPORT_PATH'] + '/lib/ui.rb'
# require ENV['TM_SUPPORT_PATH'] + "/lib/escape.rb"
# require ENV['TM_SUPPORT_PATH'] + "/lib/tm/require_cmd.rb"
# require ENV['TM_SUPPORT_PATH'] + "/lib/tm/htmloutput.rb"
# require ENV['TM_SUPPORT_PATH'] + "/lib/web_preview"

# current document
document = []
File.open("command.go", "r+") do |file|
  document = file.readlines
end

# byte offset of cursor position from the beginning of file
cursor = document[ 0, ENV['TM_LINE_NUMBER'].to_i - 1].join().length + ENV['TM_LINE_INDEX'].to_i
output = `$TM_GOCODE -f=csv -in=#{e_sh ENV['TM_FILEPATH']} autocomplete #{cursor}`

# quit if no completions found
TextMate.exit_show_tool_tip("No completions found.") if output.length == 0

# set up images for use by DIALOG
# this probably should be done only once, somehow.
icon_plist = "{ " + [ "const", "func", "package", "type", "var" ].map { |v|
  "#{v} = '#{ENV['TM_BUNDLE_SUPPORT']}/icons/#{v}.png';"
}.join(" ") + " }"
system( ENV['DIALOG'], "images", "--register", icon_plist )

# helper function to build the choice hash
def make_completion_hash(line)
  comp = line.split(",,", 3)

  match = comp[1]
  image = comp[0]

  display = " " + comp[1]
  display += comp[0] == "func" ? comp[2].gsub(/^func/, "") : " " + comp[2]

  # input : "foo(x func(func())) (z int, k int)"
  # output: "0000111111122222210001111111111110"
  def depth_at_i(sig)
    depths = Array.new
    depth = 0
    sig.chars { |ch|
      depth-=1 if ch == ")"
      depths << depth
      depth+=1 if ch == "("
    }
    return depths
  end

  # returns function arguments in the form "x arg, y arg"
  def get_f_args(sig)
    depths = depth_at_i(sig)
    pos = sig.index(")")
    while pos != nil && depths[pos] > 0
      pos = sig.index(")", pos+1)
    end
    return sig[ Range.new(sig.index("(")+1, pos - 1) ]
  end

  def split_args(sig)
    args = Array.new
    depths = depth_at_i(sig)
    start = 0
    pos = sig.index(",")

    while pos != nil
      if depths[pos] == 0
        args << sig[ Range.new(start, pos-1) ]
        start = pos+1
      end
      pos = sig.index(",", pos+1)
    end

    lastarg = sig[ Range.new(start, sig.length) ].strip
    args << lastarg unless lastarg == ""
    return args
  end

  if comp[0] == "func"
    i = 0
    insert = "(" + split_args( get_f_args( display )).map { |m| "${#{i += 1}:"+e_snp(m)+"}" }.join(", ") + ")$0"
  else
    insert = ""
  end

  return { 'match' => match, 'display' =>  display, 'insert' => insert, 'image' => image }
end

# build the list of completion choices.
hash = output.split("\n").collect { |v| make_completion_hash( v ) }
options = { :extra_chars => "_", :case_insensitive => false }

# if there is only one match, insert. no need to show the menu
if hash.length == 1
  word = ENV['TM_CURRENT_WORD'] || ""
  snippet = hash[0]["match"].gsub(/^#{Regexp.escape(word)}/, "") + hash[0]["insert"]
  #`"$DIALOG" x-insert --snippet #{e_sh snippet}`
  #full = document.join()
  #print e_sn(full[0..cursor]) + "$0" + e_sn(full[cursor+1..-1])
  TextMate.exit_insert_snippet( snippet )
else
  TextMate::UI.complete( hash , options )
end
