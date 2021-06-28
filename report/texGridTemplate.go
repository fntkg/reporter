/*
   Copyright 2016 Vastech SA (PTY) LTD

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.

   
*/

package report

const defaultGridTemplate = `
%use square brackets as golang text templating delimiters
\documentclass{report}
\usepackage[utf8]{inputenc}
\usepackage{graphicx, xcolor}
\usepackage[margin=1in]{geometry}
\graphicspath{ {images/} }
\usepackage{lipsum} % for filler text
\usepackage{fancyhdr}
\pagestyle{fancy}
\fancyhead{} % clear all header fields
\renewcommand{\headrulewidth}{0pt} % no line in header area
\fancyfoot{} % clear all footer fields
\lfoot{\colorbox{violet}{%
  \makebox[\dimexpr\linewidth-2\fboxsep][l]{\color{white}%
    Hiberus Sistemas Informáticos Zaragoza - Tfno. 902877392
    \hfill
    \thepage
  }%
}}

\title{Informe Grafana}
\author{Hiberus sistemas}

\begin{document}


\begin{titlepage}
    \centering
    \includegraphics[width=8cm]{/root/.go/bin/templates/logo.png} % also works with logo.pdf
    \vfill
    {\bfseries\Large
        Grafana Report\\
        [[.FromFormatted]]\\to\\[[.ToFormatted]]\\
        \vskip2cm
        Hiberus sistemas\\
    }    
    \vfill
    \vfill
    \vfill
\end{titlepage}

\begin{center}
[[range .Panels]][[if .IsSingleStat]]\begin{minipage}{0.3\textwidth}
\includegraphics[width=\textwidth]{image[[.Id]]}
\end{minipage}
[[else]]\par
\vspace{0.5cm}
\includegraphics[width=\textwidth]{image[[.Id]]}
\par
\vspace{0.5cm}
[[end]][[end]]
\end{center}

\end{document}
`
