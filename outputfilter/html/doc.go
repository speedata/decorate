// Package html is an output filter to generate an html chunk wrapped in a
// div to be highlighted by CSS.
// A sample output is
//    <div class="highlight"><pre><span class="nt">&lt;root&gt;</span>
//      <span class="nt">&lt;child</span> <span class="nt">/&gt;</span>
//    <span class="nt">&lt;/root&gt;</span>
//    </pre></div>
// Which requires a suitable CSS file, a CSS from pygments should be fine.
//
// Note that this package does not export any symbols, as it actively registers
// at the process package with
//     processor.RegisterOutputFilter("html", filter{})
// in the init function. filter contains a function with this signature:
//     (f filter) Render(t processor.Tokenizer) string
package html
