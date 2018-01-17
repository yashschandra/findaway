var path=require('path');

var config={
	entry: path.join(__dirname, 'templates/main.js'),
	output: {
		path: path.join(__dirname, 'static'),
		filename: 'index.js',
	},
	/*devServer: {
		inline: true,
		port: 7777
	},*/
	module: {
		loaders: [{
				test: /\.jsx?$/,
				exclude: /node_modules/,
				loader: 'babel-loader',
				query: {
					presets: ['es2015', 'react']
				}
			}
		]
	},
};

module.exports=config;
